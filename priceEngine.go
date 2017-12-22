package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type PriceCralwer interface {
	ETH() float64
	BTC() float64
	Run()
}

type BinanceCrawler struct {
	c        *http.Client
	log      *logrus.Entry
	priceMap map[string]float64
}

func NewBinanceCrawler() *BinanceCrawler {
	return &BinanceCrawler{
		c: &http.Client{
			Timeout:   3 * time.Second,
			Transport: &http.Transport{},
		},
		log:      logrus.New().WithField("service", "engine"),
		priceMap: map[string]float64{},
	}
}

func (e *BinanceCrawler) ETH() float64 {
	return e.priceMap["ETHUSDT"]
}

func (e *BinanceCrawler) BTC() float64 {
	return e.priceMap["BTCUSDT"]
}

func (e *BinanceCrawler) Run() {
	go func() {
		for {
			resp, err := e.c.Get("https://www.binance.com/api/v1/ticker/allPrices")
			if err != nil {
				e.log.Error("fail to get prices from binance. error:", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
			defer resp.Body.Close()
			var buf bytes.Buffer
			prices := []map[string]string{}
			_, _ = io.Copy(&buf, resp.Body)
			d := json.NewDecoder(&buf)
			err = d.Decode(&prices)
			if err != nil {
				e.log.Error("fail to parse prices body. error:", err.Error())
			}

			for _, priceInfo := range prices {
				p, _ := strconv.ParseFloat(priceInfo["price"], 64)
				e.priceMap[priceInfo["symbol"]] = p
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

type MarketCapCrawler struct {
	c        *http.Client
	log      *logrus.Entry
	priceMap map[string]float64
}

func NewMarketCapCrawler() *MarketCapCrawler {
	return &MarketCapCrawler{
		c: &http.Client{
			Timeout:   3 * time.Second,
			Transport: &http.Transport{},
		},
		log:      logrus.New().WithField("service", "engine"),
		priceMap: map[string]float64{},
	}
}

func (e *MarketCapCrawler) ETH() float64 {
	return e.priceMap["ETH"]
}

func (e *MarketCapCrawler) BTC() float64 {
	return e.priceMap["BTC"]
}

func (e *MarketCapCrawler) Run() {
	go func() {
		for {
			resp, err := e.c.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")
			if err != nil {
				e.log.Error("fail to get prices from marketcap. error:", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
			var buf bytes.Buffer
			prices := []map[string]string{}
			_, _ = io.Copy(&buf, resp.Body)
			resp.Body.Close()

			d := json.NewDecoder(&buf)
			err = d.Decode(&prices)
			if err != nil {
				e.log.Error("fail to parse prices body. error:", err.Error())
			}

			for _, priceInfo := range prices {
				p, _ := strconv.ParseFloat(priceInfo["price_usd"], 64)
				e.priceMap[priceInfo["symbol"]] = p
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
