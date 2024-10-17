package bankController

import (
	bank "bankingApp/bank/service"
	"bankingApp/helper"
	userController "bankingApp/user/controller"
	"bankingApp/validations"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BankParams struct {
	BankId           int    `json:"bankId,omitempty"`
	BankName         string `json:"bankName"`
	BankAbbreviation string `json:"bankAbbreviation"`
	IsActive         bool   `json:"isActive"`
}

func CreateNewBank(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateNewBankCalled")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)

	adminUser, err := userController.GetAdminObjectFromContext(r)
	if err != nil {
		panic(err)
	}

	tempBankParams := &BankParams{}
	_, err2 := helper.ParseRequestBody(r, tempBankParams)
	if err2 != nil {
		panic(err2)
	}
	errValidations := helper.ValidateAll(
		validations.ValidateIfNotEmpty("bankName", tempBankParams.BankName),
		validations.ValidateIfNotEmpty("bankAbbreviation", tempBankParams.BankAbbreviation),
	)
	if errValidations != nil {
		panic(errValidations)
	}

	tempBankInterface, err2 := adminUser.NewBank(tempBankParams.BankName, tempBankParams.BankAbbreviation)
	bank := &BankParams{
		BankId:           tempBankInterface.GetId(),
		BankName:         tempBankInterface.GetName(),
		BankAbbreviation: tempBankInterface.GetAbbreviation(),
		IsActive:         tempBankInterface.GetIsActive()}

	helper.PackRequestBody(w, http.StatusCreated, "Bank Successfully Created", bank)
}

func GetBankById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBankById Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)

	tempBankIdInStringFormat, ok := mux.Vars(r)["bankId"]
	if !ok {
		panic(errors.New("BankId Not Found"))
	}
	tempBankId, err2 := strconv.Atoi(tempBankIdInStringFormat)
	if err2 != nil {
		panic(err2)
	}

	errValidations := helper.ValidateAll(
		validations.ValidateIfNotZero("bankId", tempBankId),
		validations.ValidateIfNotNegative("bankId", tempBankId))

	if errValidations != nil {
		panic(errValidations)
	}

	tempBankInterface, err2 := bank.GetBankById(tempBankId)
	if err2 != nil {
		panic(err2)

	}
	bank := &BankParams{
		BankId:           tempBankInterface.GetId(),
		BankName:         tempBankInterface.GetName(),
		BankAbbreviation: tempBankInterface.GetAbbreviation(),
		IsActive:         tempBankInterface.GetIsActive()}

	helper.PackRequestBody(w, http.StatusCreated, "Bank Retrieved", bank)
}

func GetAllBanks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBankAllBanks Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)

	tempAllBankInterface := bank.GetAllBanks()
	var allBanks []*BankParams
	allBanks = make([]*BankParams, len(tempAllBankInterface))
	for i, tempBankInterface := range tempAllBankInterface {
		allBanks[i] = &BankParams{
			BankId:           tempBankInterface.GetId(),
			BankName:         tempBankInterface.GetName(),
			BankAbbreviation: tempBankInterface.GetAbbreviation(),
			IsActive:         tempBankInterface.GetIsActive()}

	}
	helper.PackRequestBody(w, http.StatusCreated, "All Banks Successfully Retrieved", allBanks)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteBank Called")
	defer func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.(error).Error())
		}
	}(w, r)
	adminUser, err := userController.GetAdminObjectFromContext(r)
	if err != nil {
		panic(err)
	}
	tempBankIdInStringFormat, ok := mux.Vars(r)["bankId"]
	if !ok {
		panic(errors.New("BankId Not Found"))
	}
	tempBankId, err2 := strconv.Atoi(tempBankIdInStringFormat)
	if err2 != nil {
		panic(err2)
	}
	bank, err := adminUser.DeleteBank(tempBankId)
	if err != nil {
		panic(err)
	}

	helper.PackRequestBody(w, http.StatusCreated, "All Banks Successfully Retrieved", bank)
}
