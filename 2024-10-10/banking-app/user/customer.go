package user

import (
	"bankingApp/bank"
	"bankingApp/bankAccount"
	"errors"
	"slices"
)

type Customer struct {
	totalBalance float64
	accounts     []*bankAccount.BankAccount
}

func newCustomer(customerParameters map[string]interface{}) *Customer {

	var tempEmptyAccounts []*bankAccount.BankAccount
	return &Customer{
		totalBalance: 0,
		accounts:     tempEmptyAccounts,
	}
}
func (customer *Customer) deleteCustomer() (float64, error) {
	var tempBalance float64 = 0
	for _, account := range customer.accounts {
		balance, err := account.CloseBankAccount()
		if err != nil {
			return 0, err
		}
		tempBalance += balance
	}
	return tempBalance, nil
}

func (customer *Customer) openNewBankAccount(bankId int, userId int) (*bankAccount.BankAccount, error) {
	bnk, err := bank.GetBankById(bankId)
	if err != nil {
		return nil, err
	}
	account, err1 := bnk.OpenNewBankAccount(userId)
	customer.addNewBankAccountToList(account)
	customer.addBalance(account.GetBalance())

	return account, err1
}
func (customer *Customer) closeBankAccount(bankId int, userId int, accountNumber int) error {
	bnk, err := bank.GetBankById(bankId)
	if err != nil {
		return err
	}
	leftBalance, err1 := bnk.CloseBankAccount(accountNumber, userId)
	customer.subtractBalance(leftBalance)
	customer.removeBankAccountFromList(bankId, accountNumber)

	return err1
}

func (customer *Customer) withdrawMoney(accountNumber int, bankId int, amount float64) error {
	err1 := validateAccountNumber(accountNumber)
	if err1 != nil {
		return err1
	}

	account, err := customer.getAccountByNumber(accountNumber, bankId)
	if err != nil {
		return err
	}
	err2 := account.WithdrawMoney(amount)
	if err2 != nil {
		return err2
	}
	customer.subtractBalance(amount)
	return nil

}
func (customer *Customer) depositMoney(accountNumber int, bankId int, amount float64) error {
	err1 := validateAccountNumber(accountNumber)
	if err1 != nil {
		return err1
	}

	account, err := customer.getAccountByNumber(accountNumber, bankId)
	if err != nil {
		return err
	}
	err2 := account.DepositMoney(amount)
	if err2 != nil {
		return err2
	}
	customer.addBalance(amount)
	return err2

}

func (customer *Customer) addNewBankAccountToList(account *bankAccount.BankAccount) {
	customer.accounts = append(customer.accounts, account)
}

func (customer *Customer) removeBankAccountFromList(bankId int, accountNumber int) {
	index := -1
	for i, account := range customer.accounts {
		if accountNumber == account.GetAccountNumber() && bankId == account.GetBankId() {
			index = i
		}
	}
	if index != -1 {
		customer.accounts = slices.Delete(customer.accounts, index, index+1)
	}

}

func (customer *Customer) getAccountByNumber(accountNumber int, bankId int) (*bankAccount.BankAccount, error) {
	for _, account := range customer.accounts {
		if account.GetAccountNumber() == accountNumber && account.GetBankId() == bankId {
			return account, nil
		}
	}
	return nil, errors.New("Account Not Found")
}

//updates

func (customer *Customer) addBalance(amount float64) {
	customer.totalBalance += amount
}

func (customer *Customer) subtractBalance(amount float64) {
	customer.totalBalance -= amount
}

//validations

func validateAccountNumber(accountNumber int) error {
	if accountNumber < 0 {
		return errors.New("invalid Account Number")
	}
	return nil
}
