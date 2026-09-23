package internal

import (
	"flag"
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

	c.Redirect(http.StatusFound, config.Origin)
}
