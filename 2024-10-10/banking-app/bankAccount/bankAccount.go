package bankAccount

import (
	"bankingApp/helper"
	"errors"
)

type BankAccount struct {
	accountNumber int
	balance       float64
	bankId        int
	customerId    int
	isActive      bool
}

func NewBankAccount(
	accountNumber int,
	customerId int,
	initialBalance float64,
	bankId int) (*BankAccount, error) {

	err := helper.ValidateAll(
		validateBankId(bankId),
		validateAccountNumber(accountNumber),
		validateCustomerId(customerId))
	if err != nil {
		return nil, err
	}
	tempBankAccount := &BankAccount{
		accountNumber: accountNumber,
		balance:       initialBalance,
		customerId:    customerId,
		bankId:        bankId,
		isActive:      true,
	}
	return tempBankAccount, nil
}

func (bankAccount *BankAccount) CloseBankAccount() (float64, error) {
	err := bankAccount.validateIfActive()
	if err != nil {
		return 0, err
	}
	bankAccount.isActive = false
	tempBalance := bankAccount.balance
	bankAccount.balance = 0
	return tempBalance, nil
}

func (account *BankAccount) GetBalance() float64 {
	return account.balance
}

func (account *BankAccount) GetAccountNumber() int {
	return account.accountNumber
}

func (account *BankAccount) GetCustomerId() int {
	return account.customerId
}

func (account *BankAccount) GetBankId() int {
	return account.bankId
}

func (account *BankAccount) DepositMoney(depositAmount float64) error {
	err1 := account.validateIfActive()
	if err1 != nil {
		return err1
	}
	err := validateDepositAmount(depositAmount)
	if err != nil {
		return err
	}
	account.creditToBalance(depositAmount)

	return nil

}

func (account *BankAccount) WithdrawMoney(withdrawAmount float64) error {

	err := helper.ValidateAll(
		account.validateIfActive(),
		validateWithdrawAmount(withdrawAmount),
		account.validateIfBalanceToWithdraw(withdrawAmount))
	if err != nil {
		return err
	}
	account.debitFromBalance(withdrawAmount)

	return nil

}

// private Functions

func (account *BankAccount) creditToBalance(depositAmount float64) {
	newBalance := account.GetBalance() + depositAmount
	account.balance = newBalance
}

func (account *BankAccount) debitFromBalance(withdrawalAmount float64) {
	newBalance := account.GetBalance() - withdrawalAmount
	account.balance = newBalance
}

// validate
func validateAccountNumber(accountNumber int) error {
	if accountNumber < 0 {
		return errors.New("account Number Cannot be Negative")
	}
	return nil
}

func validateCustomerId(customerId int) error {
	if customerId < 0 {
		return errors.New("customerId Cannot be Negative")
	}
	return nil
}
func validateBankId(bankId int) error {
	if bankId < 0 {
		return errors.New("bank Id Cannot be Negative")
	}
	return nil
}

func validateDepositAmount(depositAmount float64) error {
	if depositAmount < 0 {
		return errors.New("deposit Amount Cannot be Negative")
	}
	if depositAmount == 0 {
		return errors.New("deposit Amount Cannot be Zero")
	}
	return nil
}
func validateWithdrawAmount(withdrawAmount float64) error {
	if withdrawAmount < 0 {
		return errors.New("withdraw Amount Cannot be Negative")
	}
	if withdrawAmount == 0 {
		return errors.New("withdraw Amount Cannot be Zero")
	}
	return nil
}

func (account *BankAccount) validateIfBalanceToWithdraw(withdrawAmount float64) error {
	if account.GetBalance() < withdrawAmount {
		return errors.New("Insufficient Balance to withdraw")
	}
	return nil
}

func (account *BankAccount) validateIfActive() error {
	if !account.isActive {
		return errors.New("Account Donot Exist")
	}
	return nil
}
