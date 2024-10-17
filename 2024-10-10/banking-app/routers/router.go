package routers

import (
	bankController "bankingApp/bank/controller"
	bankAccountController "bankingApp/bankAccount/controller"
	"bankingApp/middleware/auth"
	userController "bankingApp/user/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterCusomterRouter(parentRouter *mux.Router) {
	adminRouter := parentRouter.NewRoute().Subrouter()
	adminRouter.Use(auth.AuthenticationMiddleware, auth.ValidateAdminPermissionsMiddleware)
	adminRouter.HandleFunc("/customers", userController.CreateNewCustomer).Methods(http.MethodPost)
	adminRouter.HandleFunc("/customers", userController.GetAllCustomers).Methods(http.MethodGet)
	adminRouter.HandleFunc("/customers/{id}", userController.GetCustomerById).Methods(http.MethodGet)
	adminRouter.HandleFunc("/customers/{id}", userController.DeleteCustomerById).Methods(http.MethodDelete)
}

func RegisterBankRouter(parentRouter *mux.Router) {
	adminRouter := parentRouter.NewRoute().Subrouter()
	authenticatedRouter := parentRouter.NewRoute().Subrouter()
	adminRouter.Use(auth.AuthenticationMiddleware, auth.ValidateAdminPermissionsMiddleware)
	authenticatedRouter.Use(auth.AuthenticationMiddleware)
	adminRouter.HandleFunc("/banks", bankController.CreateNewBank).Methods(http.MethodPost)
	authenticatedRouter.HandleFunc("/banks", bankController.GetAllBanks).Methods(http.MethodGet)
	authenticatedRouter.HandleFunc("/banks/{bankId}", bankController.GetBankById).Methods(http.MethodGet)
	adminRouter.HandleFunc("/banks/{bankId}", bankController.DeleteBank).Methods(http.MethodDelete)
}

func RegisterBankAccountRouter(parentRouter *mux.Router) {
	customerRouter := parentRouter.PathPrefix("/customers/{id}").Subrouter()
	customerRouter.Use(auth.AuthenticationMiddleware, auth.ValidateCustomerPermissionMiddleware)

	customerRouter.HandleFunc("/banks/{bankId}/accounts", bankAccountController.OpenNewBankAccountApi).Methods(http.MethodPost)
	customerRouter.HandleFunc("/accounts", bankAccountController.GetAllBankAccounts).Methods(http.MethodGet)
	customerRouter.HandleFunc("/banks/{bankId}/accounts/{accountNumber}", bankAccountController.GetAccountByAccountNumber).Methods(http.MethodGet)
	customerRouter.HandleFunc("/banks/{bankId}/accounts/{accountNumber}/withdraw", bankAccountController.WithdrawAmount).Methods(http.MethodPost)
	customerRouter.HandleFunc("/banks/{bankId}/accounts/{accountNumber}/deposit", bankAccountController.DepositMoney).Methods(http.MethodPost)
	customerRouter.HandleFunc("/banks/{bankId}/accounts/{accountNumber}/transfer", bankAccountController.TransferMoney).Methods(http.MethodPost)
	customerRouter.HandleFunc("/banks/{bankId}/accounts/{accountNumber}", bankAccountController.CloseAccount).Methods(http.MethodDelete)
	customerRouter.HandleFunc("/banks/{bankId}/accounts/{accountNumber}/transactions", bankAccountController.GetTransactions).Methods(http.MethodGet)
}
