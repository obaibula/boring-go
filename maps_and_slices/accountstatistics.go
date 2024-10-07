package maps_and_slices

import (
	"cmp"
	"github.com/shopspring/decimal"
	"slices"
	"strings"
)

type BinaryOperator[T any] func(left, right T) T

var (
	AccountLastNameComparator = func(a Account, b Account) int {
		r := cmp.Compare(a.LastName, b.LastName)
		if r == 0 {
			return cmp.Compare(a.FirstName, b.FirstName)
		}
		return r
	}
	AccountIdBinaryOperator = func(left, right Account) Account {
		if left.ID < right.ID {
			return left
		} else {
			return right
		}
	}
)

// GetTotalBalance calculates and returns the total balance from a slice of accounts.
// If the slice is empty or nil, it returns false using the comma-ok idiom.
//
// Params:
//   - accounts: A slice of Account structures from which to calculate the total balance.
//
// Returns:
//   - A decimal.Decimal representing the total balance of all accounts combined.
//   - A boolean indicating whether the slice contained any accounts (false if the slice was empty).
func GetTotalBalance(accounts []Account) (decimal.Decimal, bool) {
	total := decimal.Zero
	ok := false
	for _, a := range accounts {
		total = total.Add(a.Balance)
		ok = true
	}
	return total, ok
}

// PartitionAccountsBySex partitions a slice of accounts into a map based on their sex.
// The resulting map contains two keys: true for Male accounts and false for Female accounts.
//
// Params:
//   - accounts: A slice of Account structures to partition by sex.
//
// Returns:
//   - A map where the keys are booleans indicating the sex (true for Male, false for Female),
//     and the values are slices of Account structures corresponding to each sex.
func PartitionAccountsBySex(accounts []Account) map[bool][]Account {
	partitioned := make(map[bool][]Account)
	for _, a := range accounts {
		partitioned[a.Sex == Male] = append(partitioned[a.Sex == Male], a)
	}
	return partitioned
}

// GroupAccountsByEmailDomain groups a slice of accounts by their email domain.
// The resulting map contains email domains as keys and slices of accounts associated with each domain as values.
//
// Params:
//   - accounts: A slice of Account structures to group by email domain.
//
// Returns:
//   - A map where the keys are email domains (extracted from the account's email address), and the values
//     are slices of Account structures that share the same email domain.
func GroupAccountsByEmailDomain(accounts []Account) map[string][]Account {
	grouped := make(map[string][]Account)
	for _, a := range accounts {
		if _, domain, ok := strings.Cut(a.Email, "@"); ok {
			grouped[domain] = append(grouped[domain], a)
		}
	}
	return grouped
}

// SortByLastNameThenByFirstName sorts a slice of accounts first by last name and then by first name.
// It returns a new sorted slice of Account structures.
//
// Params:
//   - accounts: A slice of Account structures to be sorted.
//
// Returns:
//   - A new slice of Account structures sorted by last name and then by first name.
func SortByLastNameThenByFirstName(accounts []Account) []Account {
	return slices.SortedFunc(slices.Values(accounts), AccountLastNameComparator)
}

// FindRichestPerson finds and returns the account with the highest balance from a slice of accounts.
// If the slice is empty, it returns false using the comma-ok idiom.
//
// In case multiple accounts have the same highest balance, the function uses a provided mergeFunction
// to resolve the tie. If no mergeFunction is provided (i.e., it is nil), the default behavior is to
// resolve the tie by selecting the account with the lower ID, using prepared AccountIdBinaryOperator.
//
// Params:
//   - accounts: A slice of Account structures to search through.
//   - mergeFunction: A custom BinaryOperator that determines how to resolve ties
//     between accounts with equal balances. If nil, pre created AccountIdBinaryOperator is used.
//
// Returns:
// - The Account with the highest balance.
// - A boolean indicating whether the slice contained any accounts (false if the slice was empty).
func FindRichestPerson(accounts []Account, mergeFunction BinaryOperator[Account]) (Account, bool) {
	var result Account

	if mergeFunction == nil {
		mergeFunction = AccountIdBinaryOperator
	}

	if len(accounts) > 0 {
		result = accounts[0]
	}
	for i := 1; i < len(accounts); i++ {
		if accounts[i].Balance.GreaterThan(result.Balance) {
			result = accounts[i]
		} else if accounts[i].Balance.Equal(result.Balance) {
			result = mergeFunction(result, accounts[i])
		}
	}

	return result, len(accounts) > 0
}
