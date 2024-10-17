package routers

import (
	bankController "bankingApp/bank/controller"
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
	adminRouter.HandleFunc("/banks/{bankId}", bankController.GetBankById).Methods(http.MethodGet)
	adminRouter.HandleFunc("/banks/{bankId}", bankController.DeleteBank).Methods(http.MethodDelete)
}
