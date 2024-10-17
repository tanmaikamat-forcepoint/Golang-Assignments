package bankAccount

import (
	"strconv"
	"time"
)

var DEBIT_ENTRY = "debit"
var CREDIT_ENTRY = "credit"

type Passbook struct {
	accountNumber int
	transactions  []*transaction
}

type transaction struct {
	transactionId           int
	transactionType         string
	transactionAmount       float64
	transactionTimestamp    string
	balanceAfterTransaction float64
	isATransfer             bool
	otherAccountDetails     *TransferAccountDetails
	note                    string
}

type TransferAccountDetails struct {
	AccountNumber int
	BankId        int
}

//Factory

func newPassBook(accountNumber int) *Passbook {
	var tempTransactionArray []*transaction
	tempPassBook := &Passbook{
		accountNumber: accountNumber,
		transactions:  tempTransactionArray,
	}
	return tempPassBook
}

func newTransaction(
	transactionId int,
	transactionType string,
	transactionAmount float64,
	transactionTimestamp string,
	balanceAfterTransaction float64,
	isATransfer bool,
	otherAccountDetails *TransferAccountDetails,
	note string) *transaction {
	return &transaction{
		transactionId:           transactionId,
		transactionType:         transactionType,
		transactionAmount:       transactionAmount,
		transactionTimestamp:    transactionTimestamp,
		balanceAfterTransaction: balanceAfterTransaction,
		isATransfer:             isATransfer,
		otherAccountDetails:     otherAccountDetails,
		note:                    note,
	}
}

func (passbook *Passbook) addNewWithdrawToPassbook(
	amount float64,
	finalBalance float64, note string) {
	nextTransactionId := passbook.getNextTransactionId()

	tempTransaction := newTransaction(nextTransactionId, DEBIT_ENTRY, amount, time.Now().UTC().String(), finalBalance, false, nil, note)
	passbook._addTransactionToPassbook(tempTransaction)
}
func (passbook *Passbook) addNewDepositToPassbook(
	amount float64,
	finalBalance float64,
	note string) {
	nextTransactionId := passbook.getNextTransactionId()

	tempTransaction := newTransaction(nextTransactionId, CREDIT_ENTRY, amount, time.Now().UTC().String(), finalBalance, false, nil, note)
	passbook._addTransactionToPassbook(tempTransaction)
}
func (passbook *Passbook) addNewTransferToPassbook(
	amount float64,
	finalBalance float64,
	entryType string,
	otherBankId int,
	otherBankAccountNumber int, note string) int {
	nextTransactionId := passbook.getNextTransactionId()
	tempTransferAccountDetails := &TransferAccountDetails{
		AccountNumber: otherBankAccountNumber,
		BankId:        otherBankId,
	}

	tempTransaction := newTransaction(nextTransactionId, entryType, amount, time.Now().UTC().String(), finalBalance, true, tempTransferAccountDetails, note)
	passbook._addTransactionToPassbook(tempTransaction)
	return tempTransaction.transactionId
}

func (passbook *Passbook) _addTransactionToPassbook(currentTransaction *transaction) {
	passbook.transactions = append(passbook.transactions, currentTransaction)
}

func (passbook *Passbook) getNextTransactionId() int {
	return len(passbook.transactions) + 1
}

func (passbook *Passbook) getTransactionById(transactionId int) *transaction {
	return passbook.transactions[transactionId-1]
}

func (passbook *Passbook) GetAllTransactionsAsString() string {
	tempTransaction := "==============Account Number " + strconv.Itoa(passbook.accountNumber) + "==============\n"
	for _, txn := range passbook.transactions {
		tempTransaction += txn.transactionTimestamp + "\t" + txn.note + "\t" + txn.transactionType + "\t" + strconv.Itoa(int(txn.transactionAmount)) + "\t" + strconv.Itoa(int(txn.balanceAfterTransaction)) + "\n"

	}
	return tempTransaction

}
func (passbook *Passbook) GetAllTransactions() []*transaction {

	return passbook.transactions

}

func (txn *transaction) GetId() int {
	return txn.transactionId
}

func (txn *transaction) GetTransactionAmount() float64 {
	return txn.transactionAmount
}

func (txn *transaction) GetBalanceAfterTransaction() float64 {
	return txn.balanceAfterTransaction
}

func (txn *transaction) GetOtherAccountDetailsForTransfer() *TransferAccountDetails {
	return txn.otherAccountDetails
}

func (txn *transaction) GetTransactionType() string {
	return txn.transactionType
}
func (txn *transaction) GetNote() string {
	return txn.note
}
