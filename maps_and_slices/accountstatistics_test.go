package maps_and_slices

import (
	"testing"
)

func TestFindRichestPerson(t *testing.T) {
	table := []struct {
		name          string
		foundRichest  bool
		want          Account
		accounts      []Account
		mergeFunction AccountBinaryOperator
	}{
		{
			"finds richest person among multiple accounts",
			true,
			Accounts[3],
			Accounts,
			nil,
		},
		{
			"returns false when accounts list is nil",
			false,
			ZeroAccount,
			nil,
			nil,
		},
		{
			"returns false when accounts list is empty",
			false,
			ZeroAccount,
			[]Account{},
			nil,
		},
		{
			"finds account with zero balance",
			true,
			ZeroAccount,
			[]Account{ZeroAccount},
			nil,
		},
		{
			"finds account with zero balance among several zero-balance accounts",
			true,
			ZeroAccount,
			[]Account{ZeroAccount, ZeroAccount, ZeroAccount, ZeroAccount},
			nil,
		},
		{
			"finds account with negative balance",
			true,
			NegativeAccount,
			[]Account{NegativeAccount},
			nil,
		},
		{
			"finds account with negative balance among several negative-balance accounts",
			true,
			NegativeAccount,
			[]Account{NegativeAccount, NegativeAccount, NegativeAccount, NegativeAccount},
			nil,
		},
		{
			"resolves balance collision using custom merge function",
			true,
			ConflictingAccount,
			AccountsWithBalanceCollision,
			func(left, right Account) Account {
				if left.LastName < right.LastName {
					return left
				} else {
					return right
				}
			},
		},
		{
			"resolves balance collision using default merge function when mergeFunction is nil",
			true,
			ConflictingAccount,
			AccountsWithBalanceCollision,
			nil,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := FindRichestPerson(tt.accounts, tt.mergeFunction)
			if ok != tt.foundRichest {
				t.Fatalf("expected ok as %t, got %t", tt.foundRichest, ok)
			}
			if tt.want != got {
				t.Errorf("want %+v, got %+v", tt.want, got)
			}
		})
	}
}
