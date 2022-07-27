package service

import (
	"errors"
	"testbank/bank"
	"testbank/pkg/repository"

	"golang.org/x/text/currency"
)

type Account struct {
	repository repository.Account
}

func NewAccountService(r repository.Account) *Account {
	return &Account{repository: r}
}

func (s *Account) CreateAccount(userId string, a bank.Account) (int, error) {
	curr, err := checkCurrency(a.Currency)
	if err != nil && !curr {
		return 0, errors.New("This currrensy not supported")
	}
	return s.repository.CreateAccount(userId, a)
}
func (s *Account) AccountTransaction(senderId int, a bank.AccountTransaction) error {
	curr, err := checkCurrency(a.Currency)
	if err != nil && !curr {
		return errors.New("This currrensy not supported")
	}
	return s.repository.AccountTransaction(senderId, a)
}
func (s *Account) GetAllUserAccounts(id string) ([]bank.Account, error) {
	return s.repository.GetAllUserAccounts(id)
}
func (s *Account) GetTransaction(id string) ([]bank.TransactionList, error) {
	return s.repository.GetTransaction(id)
}

func checkCurrency(curr string) (bool, error) {
	switch curr {
	case currency.USD.Amount(curr).Currency().String():
		return true, nil
	case currency.MXN.Amount(curr).Currency().String():
		return true, nil
	case "COP":
		return true, nil
	}
	return false, errors.New("This currrensy not supported")
}
