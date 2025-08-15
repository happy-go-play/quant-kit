package marketmaking

import "testing"

func TestCurrencyPair_ExchangeSymbol(t *testing.T) {
	tests := []struct {
		name         string
		pair         CurrencyPair
		exchangeType ExchangeType
		want         string
		wantErr      bool
	}{
		{
			name:         "Binance symbol",
			pair:         CurrencyPair{"BTC", "USDT"},
			exchangeType: ExchangeTypeBinance,
			want:         "BTCUSDT",
			wantErr:      false,
		},
		{
			name:         "Gate symbol",
			pair:         CurrencyPair{"BTC", "USDT"},
			exchangeType: ExchangeTypeGate,
			want:         "BTC_USDT",
			wantErr:      false,
		},
		{
			name:         "Kucoin symbol",
			pair:         CurrencyPair{"BTC", "USDT"},
			exchangeType: ExchangeTypeKucoin,
			want:         "BTC-USDT",
			wantErr:      false,
		},
		{
			name:         "Unsupported exchange",
			pair:         CurrencyPair{"BTC", "USDT"},
			exchangeType: "unknown",
			want:         "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pair.exchangeSymbol(tt.exchangeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("exchangeSymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("exchangeSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
