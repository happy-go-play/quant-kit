package quantkit

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
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
