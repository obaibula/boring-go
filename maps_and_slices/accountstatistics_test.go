package maps_and_slices

import (
	"cmp"
	"maps"
	"slices"
	"testing"
)

var (
	accountIdComparator = func(a, b Account) int {
		return cmp.Compare(a.ID, b.ID)
	}
	accountSliceInAnyOrderComparator = func(left []Account, right []Account) bool {
		slices.SortFunc(left, accountIdComparator)
		slices.SortFunc(right, accountIdComparator)
		return slices.Equal(left, right)
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
			Accounts[3],
			Accounts,
			nil,
		},
		{
			"returns false when accounts slice is nil",
			false,
			ZeroAccount,
			nil,
			nil,
		},
		{
			"returns false when accounts slice is empty",
			false,
			ZeroAccount,
			[]Account{},
			nil,
		},
		{
			"finds account with zero balance",
			true,
			Accounts[9],
			Accounts[9:],
			nil,
		},
		{
			"finds account with zero balance among several zero-balance accounts",
			true,
			Accounts[9],
			Accounts[6:10],
			nil,
		},
		{
			"finds account with negative balance",
			true,
			Accounts[6],
			Accounts[6:7],
			nil,
		},
		{
			"finds account with negative balance among several negative-balance accounts",
			true,
			Accounts[7],
			Accounts[6:9],
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

func TestPartitionAccountsBySex(t *testing.T) {
	table := []struct {
		name        string
		partitioned map[bool][]Account
		accounts    []Account
	}{
		{
			"partitions accounts by sex",
			AccountsPartitionedBySex,
			Accounts,
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
