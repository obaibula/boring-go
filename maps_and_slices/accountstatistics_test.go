package maps_and_slices

import (
	"cmp"
	"github.com/shopspring/decimal"
	"maps"
	"slices"
	"testing"
)

var (
	accountIdComparator = func(a, b Account) int {
		return cmp.Compare(a.ID, b.ID)
	}
	accountSliceInAnyOrderComparator = func(a []Account, b []Account) bool {
		slices.SortFunc(a, accountIdComparator)
		slices.SortFunc(b, accountIdComparator)
		return slices.Equal(a, b)
	}
	accountLastNameBinaryOperator = func(left, right Account) Account {
		if left.LastName < right.LastName {
			return left
		} else {
			return right
		}
	}
)

func TestFindRichestPerson(t *testing.T) {
	table := []struct {
		name          string
		foundRichest  bool
		want          Account
		accounts      []Account
		mergeFunction BinaryOperator[Account]
	}{
		{
			"finds richest person among multiple accounts",
			true,
			theRichestPerson,
			accounts,
			nil,
		},
		{
			"returns false when accounts slice is nil",
			false,
			zeroAccount,
			nil,
			nil,
		},
		{
			"returns false when accounts slice is empty",
			false,
			zeroAccount,
			[]Account{},
			nil,
		},
		{
			"finds account with zero balance in single entry slice",
			true,
			zeroBalanceAccount,
			[]Account{zeroBalanceAccount},
			nil,
		},
		{
			"finds account with zero balance",
			true,
			zeroBalanceAccount,
			append(negativeBalanceAccounts, zeroBalanceAccount),
			nil,
		},
		{
			"finds account with negative balance in single entry slice",
			true,
			negativeBalanceAccounts[0],
			negativeBalanceAccounts[:1],
			nil,
		},
		{
			"finds account with negative balance among several negative-balance accounts",
			true,
			accounts[7],
			negativeBalanceAccounts,
			nil,
		},
		{
			"resolves balance collision using custom merge function",
			true,
			conflictingAccount,
			accountsWithBalanceCollision,
			accountLastNameBinaryOperator,
		},
		{
			"resolves balance collision using default merge function when mergeFunction is nil",
			true,
			conflictingAccount,
			accountsWithBalanceCollision,
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

func TestPartitionAccountsBySex(t *testing.T) {
	table := []struct {
		name        string
		partitioned map[bool][]Account
		accounts    []Account
	}{
		{
			"partitions accounts by sex",
			accountsPartitionedBySex,
			accounts,
		},
		{
			"returns empty map when accounts slice is nil",
			make(map[bool][]Account),
			nil,
		},
		{
			"returns empty map when accounts slice is empty",
			make(map[bool][]Account),
			[]Account{},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got := PartitionAccountsBySex(tt.accounts)
			if !maps.EqualFunc(got, tt.partitioned, accountSliceInAnyOrderComparator) {
				t.Errorf("want %+v, got %+v", tt.partitioned, got)
			}
		})
	}
}

func TestGroupAccountsByEmailDomain(t *testing.T) {
	t.Run("groups accounts by email domain", func(t *testing.T) {
		want := accountsGroupedByEmailDomain
		got := GroupAccountsByEmailDomain(accounts)
		if !maps.EqualFunc(got, want, accountSliceInAnyOrderComparator) {
			t.Errorf("want %+v, got %+v", want, got)
		}
	})
	t.Run("groups accounts and ignores corrupted emails", func(t *testing.T) {
		want := accountsGroupedByEmailDomain
		got := GroupAccountsByEmailDomain(append(accounts, Account{Email: "bademail"}))
		if !maps.EqualFunc(got, want, accountSliceInAnyOrderComparator) {
			t.Errorf("want %+v, got %+v", want, got)
		}
	})
}

func TestGetTotalBalance(t *testing.T) {
	t.Run("gets total balance", func(t *testing.T) {
		want := decimal.NewFromFloat(1_985_001.65)
		got, ok := GetTotalBalance(accounts)
		if !ok {
			t.Fatalf("expected ok as %t, got %t", true, ok)
		}
		if !want.Equal(got) {
			t.Errorf("want %+v, got %+v", want, got)
		}
	})
	t.Run("returns false when accounts slice is empty", func(t *testing.T) {
		_, ok := GetTotalBalance([]Account{})
		if ok {
			t.Errorf("expected ok as %t, got %t", false, ok)
		}
	})
}
