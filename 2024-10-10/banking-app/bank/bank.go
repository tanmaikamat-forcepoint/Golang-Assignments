package bank

import (
	"bankingApp/bankAccount"
	"bankingApp/helper"
	"errors"
)

var allBanks []*Bank
var bankIdCounter = 0

type Bank struct {
	bankId       int
	isActive     bool
	bankName     string
	abbreviation string
	accounts     []*bankAccount.BankAccount
}

func NewBank(bankName string, bankAbbreviation string) (*Bank, error) {
	bankAbbreviation = helper.RemoveAllLeadingAndTrailingSpaces(bankAbbreviation)
	bankName = helper.RemoveAllLeadingAndTrailingSpaces(bankName)
	err := helper.ValidateAll(
		validateBankAbbreviation(bankAbbreviation),
		validateBankName(bankName))
	if err != nil {
		return nil, err
	}
	var tempEmptyAccounts []*bankAccount.BankAccount
	tempBankObj := &Bank{
		bankId:       bankIdCounter,
		bankName:     bankName,
		abbreviation: bankAbbreviation,
		isActive:     true,
		accounts:     tempEmptyAccounts,
	}
	allBanks = append(allBanks, tempBankObj)
	bankIdCounter++
	return tempBankObj, nil
}

// functions
func (bank *Bank) OpenNewBankAccount(customerId int) (*bankAccount.BankAccount, error) {

	nextAccountNumber := bank.getNextAccountNumber()

	tempBankAccount, err := bankAccount.NewBankAccount(nextAccountNumber, customerId, 1000, bank.bankId)
	if err != nil {
		return nil, err
	}
	bank.accounts = append(bank.accounts, tempBankAccount)
	return tempBankAccount, nil

}

func (bank *Bank) CloseBankAccount(accountNumber int, customerId int) (float64, error) {
	bankAccount, err := bank.getBankAccountByNumber(accountNumber)
	if err != nil {
		return 0, err
	}
	err1 := validateIfCustomerHasAccessToBankAccount(bankAccount, customerId)
	if err1 != nil {
		return 0, err
	}

	balance, err2 := bankAccount.CloseBankAccount()

	return balance, err2
}

// getters
func GetAllBanks() []*Bank {
	var tempBankArray []*Bank
	for _, bank := range allBanks {
		if !bank.IsBankActive() {
			continue
		}
		tempBankArray = append(tempBankArray, bank)
	}
	return tempBankArray
}
func GetBankById(bankId int) (*Bank, error) {
	err := validateBankId(bankId)
	if err != nil {
		return nil, err
	}
	for _, bank := range allBanks {
		if !bank.IsBankActive() {
			continue
		}
		if bank.GetId() == bankId {
			return bank, nil
		}
	}
	return nil, errors.New("Bank Not Found")
}

func (bank *Bank) TransferMoneyFrom(accountNumberTo int, bankIdTo int, amount float64, accountNumberFrom int, bankIdFrom int, note string) error {
	bankAccount, err := bank.getBankAccountByNumber(accountNumberTo)
	if err != nil {
		return err
	}

	err2 := bankAccount.TransferMoneyFrom(amount, accountNumberFrom, bankIdFrom, note)

	return err2
}

func (bank *Bank) GetId() int {
	return bank.bankId
}

func (bank *Bank) getBankAccountByNumber(accountNumber int) (*bankAccount.BankAccount, error) {
	for _, bankAccount := range bank.getAllBankAccountsForTheBank() {
		if bankAccount.GetAccountNumber() == accountNumber {
			return bankAccount, nil
		}
	}
	return nil, errors.New("Bank Account not found")

}

func (bank *Bank) getAllBankAccountsForTheBank() []*bankAccount.BankAccount {
	return bank.accounts
}
func (b *Bank) getNextAccountNumber() int {
	if len(b.accounts) == 0 {
		return 1
	}
	return b.accounts[len(b.accounts)-1].GetAccountNumber() + 1
}

func (b *Bank) IsBankActive() bool {
	return b.isActive
}

// validations
func validateBankId(bankId int) error {
	if bankId < 0 || bankId >= bankIdCounter {
		return errors.New("Invalid BankId")
	}
	return nil
}
func validateBankName(bankName string) error {
	if len(bankName) < 2 {
		return errors.New("length of BankName should atleast be 2")
	}
	return nil
}
func validateBankAbbreviation(abbreviation string) error {
	if len(abbreviation) == 0 {
		return errors.New("Abbreviation Cannot be Empty")
	}
	return nil
}

func validateIfCustomerHasAccessToBankAccount(account *bankAccount.BankAccount, customerId int) error {
	if account.GetCustomerId() != customerId {
		return errors.New("Your are Not Authorized For this Action")
	}
	return nil
}
