package internal

import "github.com/gin-gonic/gin"

func SetupRouter() {
	r := gin.Default()

	

	_ = r.Run(":8000")
}
