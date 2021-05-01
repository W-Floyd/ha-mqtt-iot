package crypto

import (
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

var CG *coingecko.Client

func Setup() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG = coingecko.NewClient(httpClient)
}

func GetPrice(CoinName string, CurrencyName string) float32 {
	price, err := CG.SimpleSinglePrice(CoinName, CurrencyName)
	for err != nil {
		price, err = CG.SimpleSinglePrice(CoinName, CurrencyName)
		time.Sleep(2 * time.Second)
	}
	return price.MarketPrice
}
