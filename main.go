package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
}

func main() {
	crawler := NewCryptoCompareCrawler()
	crawler.Run()

	r := gin.New()
	r.Use(cors.Default())
	v1 := r.Group("v1")
	v1.GET("/price", func(c *gin.Context) {
		c.JSON(200, map[string]float64{
			"eth":  crawler.ETH(),
			"btc":  crawler.BTC(),
			"usdt": crawler.USDT(),
		})
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
