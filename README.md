
<p align="center">
  <img src="readme.png" alt="coincapV2 api golang"/>
</p>

# coincapV2 #

[![GoDoc](https://godoc.org/github.com/solipsis/coincapV2?status.svg)](https://godoc.org/github.com/solipsis/coincapV2) [![Go Report Card](https://goreportcard.com/badge/github.com/solipsis/coincapV2)](https://goreportcard.com/report/github.com/solipsis/coincapV2) 

client library for interacting with the coincap.io V2 api

## Installation ##

	go get -u github.com/solipsis/coincapV2/...
  

## Usage ##

```go
import "github.com/solipsis/coincapV2/pkg/coincap"
```
GoDoc: https://godoc.org/github.com/solipsis/coincapV2/pkg/coincap

## Official API Docs ##
https://docs.coincap.io/


## Examples ##

### Get Market Data ###
```go
client := coincap.NewClient(nil)

params := &MarketsRequest{
	ExchangeID:  "binance",
	BaseSymbol:  "ETH",
	QuoteSymbol: "BTC",
	Limit:       100,
	Offset:      0,
}
markets, timestamp, err := client.Markets(params)
```

### Get Data for an Asset ###
```go
client := coincap.NewClient(nil)

params := &AssetsRequest{
	Search: "BTC",
	Limit:  4,