package quantkit

import (
	"github.com/shopspring/decimal"
)

type PriceLevelFloat struct {
	Price  float64
	Volume float64
}

func NewPriceLevelFloat(p, v float64) *PriceLevelFloat {
	return &PriceLevelFloat{Price: p, Volume: v}
}

type PriceLevelDecimal struct {
	Price  decimal.Decimal
	Volume decimal.Decimal
}

func NewPriceLevel(p, v decimal.Decimal) *PriceLevelDecimal {
	return &PriceLevelDecimal{Price: p, Volume: v}
}

func NewPriceLevelFromString(p, v string) (*PriceLevelDecimal, error) {
	price, err := decimal.NewFromString(p)
	if err != nil {
		return nil, err
	}
	volume, err := decimal.NewFromString(v)
	if err != nil {
		return nil, err
	}

	return &PriceLevelDecimal{
		Price:  price,
		Volume: volume,
	}, nil
}

func NewPriceLevelFromFloat(price, volume float64, priceDecimals, volumeDecimals int32) *PriceLevelDecimal {
	return &PriceLevelDecimal{
		Price:  FloatToDecimal(price, priceDecimals),
		Volume: FloatToDecimal(volume, volumeDecimals),
	}
}
