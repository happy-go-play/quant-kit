package quantkit

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func FormatFloatWithDecimals(f float64, decimals int) string {
	//将一个float64的数字转化为string，并保留指定的小数位数
	return strconv.FormatFloat(f, 'f', decimals, 64)
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func FloatToDecimal(f float64, decimals int32) decimal.Decimal {
	//  strconv.FormatFloat(1.23, 'f', 4, 64)  "1.2300"
	//  decimal.NewFromFloat(1.23).Round(4)    "1.23"

	//return decimal.NewFromFloat(f).Round(decimals)
	return decimal.NewFromFloatWithExponent(f, -decimals)
}

func TruncateFloat64(f float64, decimals int) float64 {
	s := strconv.FormatFloat(f, 'f', decimals, 64)
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
