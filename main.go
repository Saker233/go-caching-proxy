package main

import (
	"github.com/Saker233/go-caching-proxy/internal"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)
func main() {

	godotenv.Load("app,env")
	internal.SetupRouter()


}
