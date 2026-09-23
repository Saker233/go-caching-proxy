package internal

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port       int
	Origin     string
	ClearCache bool
}

func HandleCMD() Config {
	var port int
	var origin string
	var clear bool

	flag.BoolVar(&clear, "clear-cache", false, "Clear Cache")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&origin, "origin", "", "Origin server URL")
	flag.Parse()

	return Config{
		Port:   port,
		Origin: origin,
		ClearCache: clear,
	}
}

func SetupServer(config Config) {
	r := gin.Default()
	r.GET("/*endpoint", func(c *gin.Context) {
		proxy(c, config)
	})
	r.Run(":" + strconv.Itoa(config.Port))
}

func proxy(c *gin.Context, config Config) {

	// we need to setup redis here first and make GET if we found it send it back to the user and make header X-cache HIT
	// if we didnt found it redirect to the original URL and send back the request from there, X-Cache MISS
	body, found := checkCache(config.Origin)

	if found {
		// HIT
		c.Header("X-Cache", "HIT")
		c.Data(http.StatusOK, "application/json", body)
		return
	}
	// MISS
	resp, err := http.Get(config.Origin)
	if err != nil {
		c.JSON(http.StatusBadRequest, errRespose(err))
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errRespose(err))
		return
	}
	// Store in redis
	err = Rdb.Set(Ctx, config.Origin, body, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	c.Header("X-Cache", "MISS")
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func errRespose(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
