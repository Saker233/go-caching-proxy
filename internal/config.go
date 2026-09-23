package internal

import (
	"flag"
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

func SetupRouter(config Config) {
	r := gin.Default()

	r.Run(":" + strconv.Itoa(config.Port))
}
