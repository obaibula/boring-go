package maps_and_slices

import (
	"github.com/shopspring/decimal"
)

type BinaryOperator[T any] func(left, right T) T

var (
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
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
}
