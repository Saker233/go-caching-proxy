package internal

import (
	"flag"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port   int
	Origin string
}

func HandleCMD() Config {
	var port int
	var origin string

	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&origin, "origin", "", "Origin server URL")
	flag.Parse()

	return Config{
		Port:   port,
		Origin: origin,
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

	resp, err := http.Get(config.Origin)
	if err != nil {
		c.JSON(http.StatusBadRequest, errRespose(err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errRespose(err))
		return
	}
	c.Header("X-Cache", "MISS")
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)

}

func errRespose(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
