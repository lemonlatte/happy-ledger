package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const CRAWLER_TIMEOUT = 10 * time.Second

type PriceCralwer interface {
	ETH() float64
	BTC() float64
	USDT() float64
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

type CryptoCompareCrawler struct {
	sync.RWMutex
	c        *http.Client
	log      *logrus.Entry
	priceMap map[string]map[string]float64
}

func NewCryptoCompareCrawler() *CryptoCompareCrawler {
	return &CryptoCompareCrawler{
		c: &http.Client{
			Timeout:   5 * time.Second,
			Transport: &http.Transport{},
		},
		log:      logrus.New().WithField("service", "engine"),
		priceMap: map[string]map[string]float64{},
	}
}

func (e *CryptoCompareCrawler) ETH() float64 {
	e.RLock()
	defer e.RUnlock()
	return e.priceMap["ETH"]["USD"]
}

func (e *CryptoCompareCrawler) BTC() float64 {
	e.RLock()
	defer e.RUnlock()
	return e.priceMap["BTC"]["USD"]
}

func (e *CryptoCompareCrawler) USDT() float64 {
	e.RLock()
	defer e.RUnlock()
	return e.priceMap["USDT"]["USD"]
}

func (e *CryptoCompareCrawler) Run() {
	go func() {
		for {
			resp, err := e.c.Get("https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,USDT&tsyms=BTC,ETH,USDT,USD")
			if err != nil {
				e.log.Error("fail to get prices from cryptocompare. error:", err.Error())
				time.Sleep(CRAWLER_TIMEOUT)
				continue
			}
			var buf bytes.Buffer
			prices := map[string]map[string]float64{}
			_, _ = io.Copy(&buf, resp.Body)
			resp.Body.Close()

			d := json.NewDecoder(&buf)
			e.log.Printf(buf.String())
			err = d.Decode(&prices)
			if err != nil {
				e.log.Error("fail to parse prices body. error:", err.Error())
			}
			e.Lock()
			e.priceMap = prices
			e.Unlock()
			time.Sleep(CRAWLER_TIMEOUT)
		}
	}()
}
