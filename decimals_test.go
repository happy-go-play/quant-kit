package quantkit

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestTruncateFloat64(t *testing.T) {
	{
		f := TruncateFloat64(1.23456789, 2)
		if f != 1.23 {
			t.Errorf("expected 1, got %v", f)
		}
	}

	{
		f := TruncateFloat64(1.23456789, 0)
		if f != 1 {
			t.Errorf("expected 1, got %v", f)
		}
	}
}

func TestFloatToDecimal(t *testing.T) {
	type testCase struct {
		input    float64
		decimals int32
		expected decimal.Decimal
	}
	list := []testCase{
		{11, 0, stringToDecimal(t, "11")},
		{11.00, 0, stringToDecimal(t, "11")},
		{0.01, 2, stringToDecimal(t, "0.01")},
		{0.011, 2, stringToDecimal(t, "0.01")},
		{0.001, 3, stringToDecimal(t, "0.001")},
		{0.001000, 3, stringToDecimal(t, "0.001")},
		{0.001000, 4, stringToDecimal(t, "0.001")},
		{0.000000, 2, stringToDecimal(t, "0")},
		{0.000000, 2, stringToDecimal(t, "0.00")},
	}
	for _, item := range list {
		input := item.input
		decimals := item.decimals
		expected := item.expected
		d := FloatToDecimal(input, decimals)

		require.Truef(t, d.Equal(expected), "expected %v, got %v for input %f, %d", expected, d, input, decimals)
	}
}

func stringToDecimal(t *testing.T, s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	require.NoError(t, err, "failed to convert string to decimal: %s", s)
	return d
}

func TestDecimals(t *testing.T) {
	type testCase struct {
		input    string
		expected int32
	}
	list := []testCase{
		{"11", 0},
		{"11.00", 2},
		{"0.01", 2},
		{"0.001", 3},
		{"0.001000", 6},
		{"0.000000", 6},
	}
	for _, item := range list {
		input := item.input
		expected := item.expected
		d, err := Decimals(item.input)
		require.NoError(t, err, "failed to get decimals for input: %s", input)
		require.Equalf(t, expected, d, "expected %v, got %v for input %s", expected, d, input)
	}
}

func TestTrimmedDecimals(t *testing.T) {
	type testCase struct {
		input    string
		expected int32
	}
	list := []testCase{
		{"11", 0},
		{"11.00", 0},
		{"0.01", 2},
		{"0.001", 3},
		{"0.001000", 3},
		{"0.000000", 0},
	}
	for _, item := range list {
		input := item.input
		expected := item.expected
		d, err := TrimmedDecimals(item.input)
		require.NoError(t, err)
		require.Equalf(t, expected, d, "expected %v, got %v for input %s", expected, d, input)
	}
}

func TestIsZeroString(t *testing.T) {
	type testCase struct {
		input    string
		expected bool
	}
	list := []testCase{
		{"0", true},
		{"0.00", true},
		{"0.000", true},
		{"0.000000", true},
		{"00000.0", true},
		{"000.00", true},
		{"1", false},
		{"1.00", false},
		{"-1", false},
		{"-0.01", false},
		{"", false},
		{" ", false},
		{"0e10", false},      // 科学计数法不支持
		{"0.000.000", false}, // 非法数字
	}

	for _, item := range list {
		input := item.input
		expected := item.expected
		result := IsZeroString(input)
		require.Equalf(t, expected, result, "expected %v, got %v for input %s", expected, result, input)
	}
}
