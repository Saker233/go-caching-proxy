package main

import (
	"github.com/Saker233/go-caching-proxy/internal"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)
func main() {
	internal.RedisClient()
	godotenv.Load("app,env")
	config := internal.HandleCMD()
	internal.SetupServer(config)
	

}
