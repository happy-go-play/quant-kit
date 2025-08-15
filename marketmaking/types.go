package marketmaking

import "fmt"

type ExchangeType string

const (
	ExchangeTypeBinance ExchangeType = "binance"
	ExchangeTypeKucoin  ExchangeType = "kucoin"
	ExchangeTypeGate    ExchangeType = "gate"
)

type CurrencyPair struct {
	baseCurrency  string // Base Currency, e.g. "BTC"
	quoteCurrency string // Quote Currency, e.g. "USDT"
}

func (c *CurrencyPair) BaseCurrency() string {
	return c.baseCurrency
}

func (c *CurrencyPair) QuoteCurrency() string {
	return c.quoteCurrency
}

func (c *CurrencyPair) exchangeSymbol(exchangeType ExchangeType) (string, error) {
	switch exchangeType {
	case ExchangeTypeBinance:
		return c.baseCurrency + c.quoteCurrency, nil // e.g. "BTCUSDT"
	case ExchangeTypeGate:
		return c.baseCurrency + "_" + c.quoteCurrency, nil // e.g. "BTC_USDT"
	case ExchangeTypeKucoin:
		return c.baseCurrency + "-" + c.quoteCurrency, nil // e.g. "BTC-USDT"
	default:
		return "", fmt.Errorf("unsupported exchange type: %s", exchangeType)
	}
}

func (c *CurrencyPair) BinanceSymbol() string {
	symbol, _ := c.exchangeSymbol(ExchangeTypeBinance)
	return symbol
}

func (c *CurrencyPair) KucoinSymbol() string {
	symbol, _ := c.exchangeSymbol(ExchangeTypeKucoin)
	return symbol
}

func (c *CurrencyPair) GateSymbol() string {
	symbol, _ := c.exchangeSymbol(ExchangeTypeGate)
	return symbol
}

func NewCurrencyPair(baseCurrency string, quoteCurrency string) *CurrencyPair {
	return &CurrencyPair{baseCurrency: baseCurrency, quoteCurrency: quoteCurrency}
}
