package userController

import (
	"bankingApp/constants"
	"bankingApp/helper"
	user "bankingApp/user/service"
	"bankingApp/validations"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserBodyParams struct {
	UserId    int    `json:"userId,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func CreateNewCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create New User Called")
	adminUser, errx := GetAdminObjectFromContext(r)
	if errx != nil {
		helper.SendErrorWithCustomMessage(w, errx.Error())
		return
	}

	output, err := helper.ParseRequestBody(r, &UserBodyParams{})
	if err != nil {
		helper.SendErrorWithCustomMessage(w, err.Error())
		return
	}
	userObj := output.Body.(*UserBodyParams)

	//validations
	err1 := helper.ValidateAll(
		validations.ValidateIfNotEmpty("username", userObj.Username),
		validations.ValidateIfNotEmpty("password", userObj.Password),
		validations.ValidateIfNotEmpty("firstName", userObj.FirstName),
		validations.ValidateIfNotEmpty("lastName", userObj.LastName))

	if err1 != nil {
		helper.SendErrorWithCustomMessage(w, err1.Error())
		return
	}
	//service call
	tempUser, err2 := adminUser.NewCustomerUserWithUsernamePassword(userObj.FirstName, userObj.LastName, userObj.Username, userObj.Password)
	if err2 != nil {
		helper.SendErrorWithCustomMessage(w, err2.Error())
		return
	}
	helper.PackRequestBody(w, http.StatusCreated, "UserSuccessfully Created", tempUser)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get CUstomer By Id")
	adminUser, errx := GetAdminObjectFromContext(r)
	if errx != nil {
		helper.SendErrorWithCustomMessage(w, errx.Error())
		return
	}

	pathId := mux.Vars(r)["id"]

	userId, err := strconv.Atoi(pathId)
	fmt.Println(userId)
	if err != nil {
		fmt.Println(err)
		helper.SendErrorWithCustomMessage(w, "Customer Id Invalid")
		return
	}
	tempUser, err := adminUser.GetUserById(userId)
	if err != nil {
		helper.SendErrorWithCustomMessage(w, err.Error())
		return
	}
	err3 := json.NewEncoder(w).Encode(tempUser)
	if err3 != nil {
		fmt.Println(err3)
		helper.SendErrorWithCustomMessage(w, "Issue in Parsing")
		return
	}
	// helper.PackRequestBody(w, http.StatusAccepted, "User Retrieved", tempUser)

}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get CUstomer By Id")
	adminUser, errx := GetAdminObjectFromContext(r)
	if errx != nil {
		helper.SendErrorWithCustomMessage(w, errx.Error())
		return
	}
	var allCustomers []*user.User
	err := adminUser.GetAllCustomers(&allCustomers)
	if err != nil {
		helper.SendErrorWithCustomMessage(w, err.Error())
		return
	}
	err3 := json.NewEncoder(w).Encode(&allCustomers)
	if err3 != nil {
		fmt.Println(err3)
		helper.SendErrorWithCustomMessage(w, "Issue in Parsing")
		return
	}
}

func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {

}

func GetAdminObjectFromContext(r *http.Request) (user.AdminInterfaceWithPass, error) {
	// return
	tempAdminObj := r.Context().Value(constants.ClaimsAdminKey).(user.AdminInterfaceWithPass)
	if tempAdminObj == nil {
		return &user.UserWithUsernamePassword{}, errors.New("Admin Validation Failed")
	}
	admin, ok := tempAdminObj.(user.AdminInterfaceWithPass)
	if !ok {
		return &user.UserWithUsernamePassword{}, errors.New("Invalid Admin")
	}

	return admin, nil

}
