package bankAccountController

import (
	bankController "bankingApp/bank/controller"
	"bankingApp/helper"
	userController "bankingApp/user/controller"
	"bankingApp/validations"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BankAccountParams struct {
	AccountNumber int     `json:"accountNumber"`
	Balance       float64 `json:"balance"`
	BankId        int     `json:"bankId"`
	CustomerId    int     `json:"customerId"`
	IsActive      bool    `json:"IsActive"`
}

type TransactionParams struct {
	TransactionId       int     `json:"transactionId"`
	TransactionType     string  `json:"transactionType"`
	Note                string  `json:"note"`
	Amount              float64 `json:"amount"`
	FinalBalance        float64 `json:"finalBalance"`
	TransferToAccountId int     `json:"toAccountId,omitempty"`
	TransferToBankId    int     `json:"toBankId,omitempty"`
}

func OpenNewBankAccountApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Open New Bank Account Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)

	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}
	tempBankIdInStringFormat, ok := mux.Vars(r)["bankId"]
	if !ok {
		panic(errors.New("BankId Not Found"))
	}
	tempBankId, err2 := strconv.Atoi(tempBankIdInStringFormat)
	if err2 != nil {
		panic(err2)
	}

	tempAccountInterface, err2 := currentCustomerObj.OpenNewBankAccount(tempBankId)
	bank := &BankAccountParams{
		AccountNumber: tempAccountInterface.GetAccountNumber(),
		BankId:        tempAccountInterface.GetBankId(),
		CustomerId:    tempAccountInterface.GetCustomerId(),
		Balance:       tempAccountInterface.GetBalance(),
		IsActive:      true,
	}

	helper.PackRequestBody(w, http.StatusCreated, "Bank Successfully Created", bank)
}

func GetAccountByAccountNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Account by Number Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}

	tempBankId := bankController.ValidateAndGetBankIdFromPathParams(r)
	tempAccountNumber := ValidateAndGetAccountNumberFromPathParams(r)
	tempAccountInterface, err2 := currentCustomerObj.GetAccountByBankIdAccountNumber(tempBankId, tempAccountNumber)
	if err2 != nil {
		panic(err2)

	}
	bank := &BankAccountParams{
		AccountNumber: tempAccountInterface.GetAccountNumber(),
		BankId:        tempAccountInterface.GetBankId(),
		CustomerId:    tempAccountInterface.GetCustomerId(),
		Balance:       tempAccountInterface.GetBalance(),
		IsActive:      true,
	}

	helper.PackRequestBody(w, http.StatusCreated, "Bank Retrieved", bank)
}

func CloseAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Close Account by Number Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}

	tempBankId := bankController.ValidateAndGetBankIdFromPathParams(r)
	tempAccountNumber := ValidateAndGetAccountNumberFromPathParams(r)
	err2 := currentCustomerObj.CloseBankAccount(tempBankId, tempAccountNumber)
	if err2 != nil {
		panic(err2)

	}

	helper.PackRequestBody(w, http.StatusCreated, "Account Successfully Deleted", nil)
}

func GetAllBankAccounts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBankAllBankAccounts Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)

	user, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}
	accounts := user.GetUserBankAccounts()
	var allAccounts []*BankAccountParams
	allAccounts = make([]*BankAccountParams, len(accounts))
	for i, tempAccountInterface := range accounts {
		allAccounts[i] = &BankAccountParams{
			AccountNumber: tempAccountInterface.GetAccountNumber(),
			BankId:        tempAccountInterface.GetBankId(),
			CustomerId:    tempAccountInterface.GetCustomerId(),
			Balance:       tempAccountInterface.GetBalance(),
			IsActive:      true,
		}
	}
	helper.PackRequestBody(w, http.StatusCreated, "All Banks Successfully Retrieved", allAccounts)
}

func ValidateAndGetAccountNumberFromPathParams(r *http.Request) int {

	tempAccountNumberInStringFormat, ok := mux.Vars(r)["accountNumber"]
	if !ok {
		panic(errors.New("AccountNumber Not Found in Path"))
	}
	tempAccountNumber, err2 := strconv.Atoi(tempAccountNumberInStringFormat)
	if err2 != nil {
		panic(err2)
	}

	errValidations := helper.ValidateAll(
		validations.ValidateIfNotZero("accountNumber", tempAccountNumber),
		validations.ValidateIfNotNegative("accountNumber", tempAccountNumber))

	if errValidations != nil {
		panic(errValidations)
	}
	return tempAccountNumber

}

func WithdrawAmount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Withdraw Amount Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}

	tempBankId := bankController.ValidateAndGetBankIdFromPathParams(r)
	tempAccountNumber := ValidateAndGetAccountNumberFromPathParams(r)
	tempTransactionParams := &TransactionParams{}
	_, err2 := helper.ParseRequestBody(r, tempTransactionParams)
	if err2 != nil {
		panic(err2)
	}
	err = helper.ValidateAll(
		validations.ValidateIfNotNegative("amount", int(tempTransactionParams.Amount)),
		validations.ValidateIfNotZero("amount", int(tempTransactionParams.Amount)),
	)
	if err != nil {
		panic(err)
	}
	err3 := currentCustomerObj.WithdrawMoney(tempAccountNumber, tempBankId, tempTransactionParams.Amount)
	if err3 != nil {
		panic(err3)

	}
	// bank := &BankAccountParams{
	// 	AccountNumber: tempAccountInterface.GetAccountNumber(),
	// 	BankId:        tempAccountInterface.GetBankId(),
	// 	CustomerId:    tempAccountInterface.GetCustomerId(),
	// 	Balance:       tempAccountInterface.GetBalance(),
	// 	IsActive:      true,
	// }

	helper.PackRequestBody(w, http.StatusCreated, "Transaction Successful", nil)
}

func DepositMoney(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Withdraw Amount Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}

	tempBankId := bankController.ValidateAndGetBankIdFromPathParams(r)
	tempAccountNumber := ValidateAndGetAccountNumberFromPathParams(r)
	tempTransactionParams := &TransactionParams{}
	_, err2 := helper.ParseRequestBody(r, tempTransactionParams)
	if err2 != nil {
		panic(err2)
	}
	err = helper.ValidateAll(
		validations.ValidateIfNotNegative("amount", int(tempTransactionParams.Amount)),
		validations.ValidateIfNotZero("amount", int(tempTransactionParams.Amount)),
	)
	if err != nil {
		panic(err)
	}

	err3 := currentCustomerObj.DepositMoney(tempAccountNumber, tempBankId, tempTransactionParams.Amount)
	if err3 != nil {
		panic(err3)

	}
	// bank := &BankAccountParams{
	// 	AccountNumber: tempAccountInterface.GetAccountNumber(),
	// 	BankId:        tempAccountInterface.GetBankId(),
	// 	CustomerId:    tempAccountInterface.GetCustomerId(),
	// 	Balance:       tempAccountInterface.GetBalance(),
	// 	IsActive:      true,
	// }

	helper.PackRequestBody(w, http.StatusCreated, "Transaction Successful", nil)
}

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Withdraw Amount Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}

	tempBankId := bankController.ValidateAndGetBankIdFromPathParams(r)
	tempAccountNumber := ValidateAndGetAccountNumberFromPathParams(r)
	tempTransactionParams := &TransactionParams{}
	_, err2 := helper.ParseRequestBody(r, tempTransactionParams)
	if err2 != nil {
		panic(err2)
	}
	err = helper.ValidateAll(
		validations.ValidateIfNotNegative("amount", int(tempTransactionParams.Amount)),
		validations.ValidateIfNotZero("amount", int(tempTransactionParams.Amount)),
		validations.ValidateIfNotZero("accountNumber", tempTransactionParams.TransferToAccountId),
		validations.ValidateIfNotZero("bankId", tempTransactionParams.TransferToBankId))

	if err != nil {
		panic(err)
	}

	err3 := currentCustomerObj.TransferMoneyTo(tempAccountNumber, tempBankId, tempTransactionParams.Amount, tempTransactionParams.TransferToAccountId, tempTransactionParams.TransferToBankId, "Transfer")
	if err3 != nil {
		panic(err3)

	}
	// bank := &BankAccountParams{
	// 	AccountNumber: tempAccountInterface.GetAccountNumber(),
	// 	BankId:        tempAccountInterface.GetBankId(),
	// 	CustomerId:    tempAccountInterface.GetCustomerId(),
	// 	Balance:       tempAccountInterface.GetBalance(),
	// 	IsActive:      true,
	// }

	helper.PackRequestBody(w, http.StatusCreated, "Transaction Successful", nil)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Account by Number Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	currentCustomerObj, err := userController.GetCusomterObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	pathId, ok := mux.Vars(r)["id"]
	if !ok {
		panic(errors.New("CustomerId Not Found In Path"))
	}

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		panic(errors.New("Invalid Customer Id"))
	}
	if userId != currentCustomerObj.GetUserId() {
		panic(errors.New("Unauthorized Access"))

	}

	tempBankId := bankController.ValidateAndGetBankIdFromPathParams(r)
	tempAccountNumber := ValidateAndGetAccountNumberFromPathParams(r)
	tempAccountInterface, err2 := currentCustomerObj.GetAccountByBankIdAccountNumber(tempBankId, tempAccountNumber)
	if err2 != nil {
		panic(err2)

	}

	alltransactionSlice := tempAccountInterface.GetPassbook().GetAllTransactions()
	var allTransactions []*TransactionParams
	allTransactions = make([]*TransactionParams, len(alltransactionSlice))
	for i, txn := range alltransactionSlice {

		tempOtherAccountNumber := 0
		tempOtherBankId := 0

		if txn.GetOtherAccountDetailsForTransfer() != nil {
			tempOtherAccountNumber = txn.GetOtherAccountDetailsForTransfer().AccountNumber
			tempOtherBankId = txn.GetOtherAccountDetailsForTransfer().BankId
		}
		allTransactions[i] = &TransactionParams{
			TransactionId:       txn.GetId(),
			Amount:              txn.GetTransactionAmount(),
			TransferToAccountId: tempOtherAccountNumber,
			TransferToBankId:    tempOtherBankId,
			Note:                txn.GetNote(),
			TransactionType:     txn.GetTransactionType(),
			FinalBalance:        txn.GetBalanceAfterTransaction(),
		}
	}
	helper.PackRequestBody(w, http.StatusCreated, "Transactions Retrieved", allTransactions)
}
