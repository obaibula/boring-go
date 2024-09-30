package maps_and_slices

// PartitionAccountsBySex partitions a slice of Accounts into a map based on their sex.
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

// GroupAccountsByEmailDomain groups a slice of Accounts by their email domain.
// The resulting map contains email domains as keys and slices of Accounts associated with each domain as values.
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

// FindRichestPerson finds and returns the account with the highest balance from a slice of Accounts.
// If the slice is empty, it returns false using the comma-ok idiom.
//
// In case multiple accounts have the same highest balance, the function uses a provided mergeFunction
// to resolve the tie. If no mergeFunction is provided (i.e., it is nil), the default behavior is to
// resolve the tie by selecting the account with the lower ID, using prepared AccountIdBinaryOperator.
//
// Params:
//   - accounts: A slice of Account structures to search through.
//   - mergeFunction: A custom BinaryOperator that determines how to resolve ties
//     between accounts with equal balances. If nil, AccountIdBinaryOperator is used.
//
// Returns:
// - The Account with the highest balance.
// - A boolean indicating whether the slice contained any accounts (false if the slice was empty).
func FindRichestPerson(accounts []Account, mergeFunction BinaryOperator[Account]) (Account, bool) {
	panic("not implemented")
}

type BinaryOperator[T any] func(left, right T) T

func AccountIdBinaryOperator(left, right Account) Account {
	if left.ID < right.ID {
		return left
	} else {
		return right
	}
}
