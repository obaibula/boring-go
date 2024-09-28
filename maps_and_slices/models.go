package maps_and_slices

import (
	"github.com/shopspring/decimal"
	"time"
)

type Account struct {
	ID           int64
	FirstName    string
	LastName     string
	Email        string
	Birthday     time.Time
	Sex          Sex
	CreationDate time.Time
	Balance      decimal.Decimal
}

type Sex int

const (
	Male = iota
	Female
)
