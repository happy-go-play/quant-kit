package marketmaking

import "github.com/shopspring/decimal"

type TradeStat struct {
	CurrencyPair CurrencyPair

	SellVolume      decimal.Decimal // Base Currency Sell Volume
	SellQuoteVolume decimal.Decimal
	SellAvgPrice    decimal.Decimal // Base Currency Sell Average Price

	BuyVolume      decimal.Decimal // Base Currency Buy Volume
	BuyQuoteVolume decimal.Decimal
	BuyAvgPrice    decimal.Decimal // Base Currency Buy Average Price
}
