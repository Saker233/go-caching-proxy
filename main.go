package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Saker233/go-caching-proxy/internal"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	godotenv.Load("app.env")

	config := internal.HandleCMD()
	if config.ClearCache {
		internal.RedisClient()

		err := internal.Rdb.FlushDB(internal.Ctx).Err()
		if err != nil {
			log.Fatal("Failed to flusgh DB", err)
		}
		return
	}
	internal.RedisClient()
	go internal.SetupServer(config)
	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/products", config.Port))
	if err != nil {
		log.Fatal("Failed to call proxy:", err)
	}

	defer resp.Body.Close()
	log.Println("Proxy response:", resp.Header)
}
