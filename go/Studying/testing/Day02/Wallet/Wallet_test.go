package wallet

import (
	"math"
	"reflect"
	"testing"
)

func TestWallet_Deposit(t *testing.T) {
	type args struct {
		num Bitcoin
	}
	tests := []struct {
		name string
		w    *Wallet
		args args
		want Bitcoin
	}{
		// TODO: Add test cases.
		{"Wallet_Deposit", &Wallet{20}, args{44}, Bitcoin(64)},
		{"Wallet_Deposit", &Wallet{10}, args{44}, Bitcoin(54)},
		{"Wallet_Deposit", &Wallet{0}, args{0}, Bitcoin(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Deposit(tt.args.num)
			got := tt.w.Balance()
			if got != tt.want {
				t.Errorf("got %.5f want %.5f", got, tt.want)
			}
		})
	}
}

func TestWallet_Balance(t *testing.T) {
	tests := []struct {
		name string
		w    *Wallet
		want Bitcoin
	}{
		// TODO: Add test cases.
		{"Wallet_Balance", &Wallet{20}, Bitcoin(20)},
		{"Wallet_Balance", &Wallet{10}, Bitcoin(10)},
		{"Wallet_Balance", &Wallet{1}, Bitcoin(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Balance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallet.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitcoin_String(t *testing.T) {
	tests := []struct {
		name string
		b    Bitcoin
		want string
	}{
		// TODO: Add test cases.
		{"Bitcoin_String", Bitcoin(333), "333.00000 BTC"},
		{"Bitcoin_String", Bitcoin(123456789), "123456789.00000 BTC"},
		{"Bitcoin_String", Bitcoin(-5555), "-5555.00000 BTC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Bitcoin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	tests := []struct {
		name string
		w    *Wallet
		num  Bitcoin
		want Bitcoin
	}{
		// TODO: Add test cases.
		{"Wallet_Withdraw", &Wallet{0}, Bitcoin(0), Bitcoin(0)},
		{"Wallet_Withdraw", &Wallet{3}, Bitcoin(4), Bitcoin(3)},
		{"Wallet_Withdraw", &Wallet{5000}, Bitcoin(4999), Bitcoin(1)},
		{"Wallet_Withdraw", &Wallet{6e+10}, Bitcoin(5e+10), Bitcoin(1e+10)},
		{"Wallet_Withdraw", &Wallet{6.55644}, Bitcoin(5.44434), Bitcoin(1.1121)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Withdraw(tt.num)
			got := tt.w.Balance()
			// 此处涉及float64精度问题
			if math.Dim(float64(got), float64(tt.want)) > 0.000001 {
				t.Errorf("Wallet.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}
