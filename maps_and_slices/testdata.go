package maps_and_slices

import (
	"github.com/shopspring/decimal"
	"time"
)

var Accounts = []Account{
	{
		ID:           1,
		FirstName:    "Emmanuel",
		LastName:     "Macron",
		Email:        "emmanuel.macron@gmail.com",
		Birthday:     parseDate("1977-12-21"),
		Sex:          Male,
		CreationDate: parseDate("2018-05-12"),
		Balance:      decimal.NewFromFloat(520000.75),
	},
	{
		ID:           2,
		FirstName:    "Angela",
		LastName:     "Merkel",
		Email:        "angela.merkel@icloud.com",
		Birthday:     parseDate("1954-07-17"),
		Sex:          Female,
		CreationDate: parseDate("2013-11-19"),
		Balance:      decimal.NewFromFloat(610000.00),
	},
	{
		ID:           3,
		FirstName:    "Boris",
		LastName:     "Johnson",
		Email:        "boris.johnson@yahoo.com",
		Birthday:     parseDate("1964-06-19"),
		Sex:          Male,
		CreationDate: parseDate("2020-01-22"),
		Balance:      decimal.NewFromFloat(450000.25),
	},
	{
		ID:           4,
		FirstName:    "Pedro",
		LastName:     "Sánchez",
		Email:        "pedro.sanchez@outlook.com",
		Birthday:     parseDate("1972-02-29"),
		Sex:          Male,
		CreationDate: parseDate("2015-08-30"),
		Balance:      decimal.NewFromFloat(700000.30),
	},
	{
		ID:           5,
		FirstName:    "Mateusz",
		LastName:     "Morawiecki",
		Email:        "mateusz.morawiecki@hotmail.com",
		Birthday:     parseDate("1968-06-20"),
		Sex:          Male,
		CreationDate: parseDate("2014-07-15"),
		Balance:      decimal.NewFromFloat(330000.85),
	},
	{
		ID:           6,
		FirstName:    "Giuseppe",
		LastName:     "Conte",
		Email:        "giuseppe.conte@gmail.com",
		Birthday:     parseDate("1964-08-08"),
		Sex:          Male,
		CreationDate: parseDate("2019-03-05"),
		Balance:      decimal.NewFromFloat(580000.90),
	},
	{
		ID:           7,
		FirstName:    "Sebastian",
		LastName:     "Kurz",
		Email:        "sebastian.kurz@protonmail.com",
		Birthday:     parseDate("1986-08-27"),
		Sex:          Male,
		CreationDate: parseDate("2021-06-22"),
		Balance:      decimal.NewFromFloat(640000.15),
	},
	{
		ID:           8,
		FirstName:    "Ursula",
		LastName:     "von der Leyen",
		Email:        "ursula.von@gmail.com",
		Birthday:     parseDate("1958-10-08"),
		Sex:          Female,
		CreationDate: parseDate("2012-02-19"),
		Balance:      decimal.NewFromFloat(570000.50),
	},
	{
		ID:           9,
		FirstName:    "Sanna",
		LastName:     "Marin",
		Email:        "sanna.marin@icloud.com",
		Birthday:     parseDate("1985-11-16"),
		Sex:          Female,
		CreationDate: parseDate("2020-11-03"),
		Balance:      decimal.NewFromFloat(495000.75),
	},
	{
		ID:           10,
		FirstName:    "Kaja",
		LastName:     "Kallas",
		Email:        "kaja.kallas@gmail.com",
		Birthday:     parseDate("1977-06-18"),
		Sex:          Female,
		CreationDate: parseDate("2016-04-28"),
		Balance:      decimal.NewFromFloat(450000.00),
	},
	ZeroAccount,
}

var ConflictingAccount = Account{
	ID:           3,
	FirstName:    "Pedro",
	LastName:     "Ránchez",
	Email:        "pedro.ranchez@outlook.com",
	Birthday:     parseDate("1972-02-29"),
	Sex:          Male,
	CreationDate: parseDate("2015-08-29"),
	Balance:      decimal.NewFromFloat(700000.30),
}

var (
	AccountsWithBalanceCollision = append(Accounts, ConflictingAccount)
	ZeroAccount                  = Account{}
	NegativeAccount              = Account{Balance: decimal.NewFromInt(-100)}
)

func parseDate(dateStr string) time.Time {
	t, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		panic(err)
	}
	return t
}
