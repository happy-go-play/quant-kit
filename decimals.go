package quantkit

import (
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
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

// Decimals 获取小数位数
//
//	"11" -> 0
//	"11.00" -> 2
//	"0.01" -> 2
//	"0.001" -> 3
//	"0.001000" -> 6
//
// 纯字符串版本会比 decimal.NewFromString 快很多。
func Decimals(s string) (int32, error) {
	// 验证是合法数字
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return 0, err
	}

	if !strings.Contains(s, ".") {
		return 0, nil
	}

	parts := strings.SplitN(s, ".", 2)
	return int32(len(parts[1])), nil
}

// TrimmedDecimals 获取小数位数，会去掉小数位末尾多余的0
//
//	"11" -> 0
//	"11.00" -> 0
//	"0.01" -> 2
//	"0.001" -> 3
//	"0.001000" -> 3
//
// 纯字符串版本会比 decimal.NewFromString 快很多。
func TrimmedDecimals(s string) (int32, error) {
	// 验证是合法数字
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return 0, err
	}

	s = strings.TrimRight(s, "0")

	if !strings.Contains(s, ".") {
		return 0, nil
	}

	parts := strings.SplitN(s, ".", 2)
	return int32(len(parts[1])), nil
}
