package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// Anthony is teaching structuring the microservices, therefore he has not written the business logic here
	// getting the actual price for tickers, he is just returning results from the map var declared

	// It is very important to not use any user declared types or JSON representations here. Business logic needs
	// be clean and only handle the business logic
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200,
	"GG":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// mimic the HTTP roundtrip for fetching the price
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
