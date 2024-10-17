package user

import (
	bank "bankingApp/bank/service"
	bankAccount "bankingApp/bankAccount/service"
	"bankingApp/helper"
	"errors"
	"fmt"
)

var allUsers []*User
var allStaff []StaffInterface
var userIdCounter = 0

type AdminInterface interface {
	NewCustomerUser(firstName, lastName string, customerParameters ...interface{}) (*User, error)
	DeleteCustomer(userId int) (float64, error)
	GetUserById(userId int) (*User, error)
	NewBank(bankName string, abbrevation string) (bank.BankInterface, error)
}

type AdminInterfaceWithPass interface {
	AdminInterface
	GetAllCustomers(customers *[]*User) error
	DeleteBank(bankId int) (bank.BankInterface, error)
	NewCustomerUserWithUsernamePassword(firstName, lastName, username, password string, customerParameters ...interface{}) (StaffInterface, error)
}
type StaffInterface interface {
	OpenNewBankAccount(bankId int) (bankAccount.BankAccountInterface, error)
	CloseBankAccount(bankId int, accountNumber int) error
	DepositMoney(accountNumber int, bankId int, amount float64) error
	WithdrawMoney(accountNumber int, bankId int, amount float64) error
	TransferMoneyTo(accountNumberFrom int, bankIdFrom int, amount float64, accountNumberTo int, bankIdTo int, note string) error
	GetTotalBalance() (float64, error)
	GetFullName() string
	GetUserId() int
	GetUserBankAccounts() []bankAccount.BankAccountInterface
	GetAccountByBankIdAccountNumber(bankId, accountNumber int) (bankAccount.BankAccountInterface, error)
}

type User struct {
	UserId    int       `json:"userId"`
	IsAdmin   bool      `json:"isAdmin"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	customer  *Customer `json:"-"`
}

func NewAdminUser(firstName, lastName string) (AdminInterface, error) {
	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempUserObject := &User{
		UserId:    userIdCounter,
		IsAdmin:   true,
		FirstName: firstName,
		LastName:  lastName,
		customer:  nil,
	}
	userIdCounter++
	allUsers = append(allUsers, tempUserObject)
	return tempUserObject, nil
}

// Admin Function
func (user *User) NewCustomerUser(firstName, lastName string, customerParameters ...interface{}) (*User, error) {

	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateIfAdmin(user),
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempUserObject := &User{
		UserId:    userIdCounter,
		IsAdmin:   false,
		FirstName: firstName,
		LastName:  lastName,
		customer:  newCustomer(customerParameters),
	}
	allStaff = append(allStaff, tempUserObject)
	userIdCounter++
	return tempUserObject, nil
}

//admin functions

// Admin Function
func (user *User) NewBank(bankName string, abbrevation string) (bank.BankInterface, error) {
	err := validateIfAdmin(user)
	if err != nil {
		return nil, err
	}
	return bank.NewBank(bankName, abbrevation)
}

// Admin Function
func (user *User) DeleteCustomer(userId int) (float64, error) {
	err := validateIfAdmin(user)
	if err != nil {
		return 0, err
	}
	customerToDelete := allUsers[userId]
	err2 := validateIfCustomer(customerToDelete)
	if err2 != nil {
		return 0, err2
	}

	return customerToDelete.customer.deleteCustomer()
}

// general
func GetAllBanks() []bank.BankInterface {
	return bank.GetAllBanks()
}

// customer Functions
// staff function
func (user *User) OpenNewBankAccount(bankId int) (bankAccount.BankAccountInterface, error) {
	err := validateIfCustomer(user)
	if err != nil {
		return nil, err
	}
	return user.customer.openNewBankAccount(bankId, user.UserId)
}

// staff function
func (user *User) CloseBankAccount(bankId int, accountNumber int) error {
	err := validateIfCustomer(user)
	if err != nil {
		return err
	}
	return user.customer.closeBankAccount(bankId, user.UserId, accountNumber)
}

func (user *User) WithdrawMoney(accountNumber int, bankId int, amount float64) error {
	err := validateIfCustomer(user)
	if err != nil {
		return err
	}
	return user.customer.withdrawMoney(accountNumber, bankId, amount)
}
func (user *User) DepositMoney(accountNumber int, bankId int, amount float64) error {
	err := validateIfCustomer(user)
	if err != nil {
		return err
	}
	return user.customer.depositMoney(accountNumber, bankId, amount)

}

func (user *User) TransferMoneyTo(accountNumberFrom int, bankIdFrom int, amount float64, accountNumberTo int, bankIdTo int, note string) error {
	err := validateIfCustomer(user)

	if err != nil {
		return err
	}
	return user.customer.transferMoney(accountNumberFrom, bankIdFrom, amount, accountNumberTo, bankIdTo, note)
}

func (user *User) GetTotalBalance() (float64, error) {
	err := validateIfCustomer(user)
	if err != nil {
		return 0, err
	}
	return user.customer.getBalance(), nil
}

// validations

func validateInputCustomerId(userId int) error {
	if userId < 0 || userIdCounter <= userId {
		return errors.New("Invalid User Id")
	}
	return nil
}
func validateFirstName(firstName string) error {
	if len(firstName) < 2 {
		return errors.New("length of First Name should atleast be 2")
	}
	return nil
}
func validateLastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("last Name Cannot be Empty")
	}
	return nil
}

func validateIfAdmin(user *User) error {

	if !user.isUserAdmin() {
		return errors.New("unauthoried Access")
	}
	return nil
}

// TODO : isStaff ->SRP
func validateIfCustomer(user *User) error {

	if user.isUserAdmin() {
		return errors.New("Only Customers can access this")
	}
	return nil
}

// getters
func (user *User) GetUserId() int {
	return user.UserId
}
func (user *User) isUserAdmin() bool {
	return user.IsAdmin
}
func (user *User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) GetUserById(userId int) (*User, error) {
	err := helper.ValidateAll(
		validateIfAdmin(user),
		validateInputCustomerId(userId))
	if err != nil {
		return nil, err
	}

	return allUsers[userId], nil

}

type UserWithUsernamePassword struct {
	User
	Username string `json:"username"`
	Password string `json:"-"`
}

var allUsersWithPass = make(map[int]*UserWithUsernamePassword)

func NewAdminUserWithIdPass(firstName, lastName, username, password string) (AdminInterfaceWithPass, error) {
	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName), validatePassword(password))
	if err != nil {
		return nil, err
	}

	hashedPassword := helper.HashPassword(password)
	tempUserObject := &UserWithUsernamePassword{
		User: User{
			UserId:    userIdCounter,
			IsAdmin:   true,
			FirstName: firstName,
			LastName:  lastName,
			customer:  nil,
		},
		Username: username,
		Password: hashedPassword,
	}
	allUsersWithPass[userIdCounter] = tempUserObject
	userIdCounter++
	return tempUserObject, nil
}

func (user *UserWithUsernamePassword) NewCustomerUserWithUsernamePassword(firstName, lastName, username, password string, customerParameters ...interface{}) (StaffInterface, error) {

	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateIfAdmin(&user.User),
		validateFirstName(firstName),
		validateLastName(lastName),
		validatePassword(password))
	if err != nil {
		return nil, err
	}
	hashedPassword := helper.HashPassword(password)
	tempUserObject := &UserWithUsernamePassword{
		User: User{
			UserId:    userIdCounter,
			IsAdmin:   false,
			FirstName: firstName,
			LastName:  lastName,
			customer:  newCustomer(customerParameters)},
		Username: username,
		Password: hashedPassword}

	allUsersWithPass[userIdCounter] = tempUserObject
	userIdCounter++
	fmt.Println(userIdCounter)
	return tempUserObject, nil
}

func (user *UserWithUsernamePassword) GetUserById(userId int) (*User, error) {
	err := helper.ValidateAll(
		validateIfAdmin(&user.User),
		validateInputCustomerId(userId))
	if err != nil {
		return nil, err
	}
	userObj := &allUsersWithPass[userId].User

	return userObj, nil

}

func (user *UserWithUsernamePassword) GetAllCustomers(customers *[]*User) error {
	err := helper.ValidateAll(
		validateIfAdmin(&user.User))

	if err != nil {
		return err
	}
	for _, user := range allUsersWithPass {
		if !user.isUserAdmin() {
			*customers = append(*customers, &user.User)
		}
	}
	return nil
}

func AuthenticateUser(userId int, password string) *User {
	err := helper.ValidateAll(
		validatePassword(password))

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(userId)
	tempUser, ok := allUsersWithPass[userId]
	if !ok {
		fmt.Println("UserName Not found1")
		return nil
	}
	hashedPassword := helper.HashPassword(password)
	fmt.Println(tempUser.getHashedPassword(), hashedPassword)
	if helper.CheckHashWithPassword(password, tempUser.getHashedPassword()) {
		return tempUser.GetUserData()
	}
	return nil

}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password should be atleast of Length 8")
	}
	return nil
}

// func GetUserForAuthById(userId int) (*User, error) {
// 	err := helper.ValidateAll(
// 		validateInputCustomerId(userId))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return allUsersWithPass[userId], nil

// }

func getUserByUsername(username string) (*UserWithUsernamePassword, error) {
	for _, user := range allUsersWithPass {
		fmt.Println(user)
		if user.Username == username {
			return user, nil
		}

	}
	return nil, errors.New("User Not Found")

}

func (user *UserWithUsernamePassword) getHashedPassword() string {
	return user.Password
}

func (user *UserWithUsernamePassword) GetUserData() *User {
	return &user.User
}

func GetAdminInterfaceWithPassById(userId int) (AdminInterfaceWithPass, error) {
	tempUser, ok := allUsersWithPass[userId]
	if !ok {
		return nil, errors.New("Admin Not Found")
	}
	if !tempUser.isUserAdmin() {
		return nil, errors.New("User is not a Admin")
	}

	return tempUser, nil
}

func (user *UserWithUsernamePassword) DeleteBank(bankId int) (bank.BankInterface, error) {
	err := helper.ValidateAll(
		validateIfAdmin(&user.User),
	)
	if err != nil {
		return nil, err
	}
	tempBankInterface, err := bank.GetBankById(bankId)
	if err != nil {
		return nil, err
	}

	err = tempBankInterface.DeleteSelf()
	if err != nil {
		return nil, err
	}
	return tempBankInterface, nil

}
func GetStaffInterfaceWithPassById(userId int) (StaffInterface, error) {
	tempUser, ok := allUsersWithPass[userId]
	if !ok {
		return nil, errors.New("User Not Found")
	}
	if tempUser.isUserAdmin() {
		return nil, errors.New("User is not a Customer")
	}

	return tempUser, nil
}

func (user *User) GetUserBankAccounts() []bankAccount.BankAccountInterface {
	return user.customer.accounts
}

func (user *User) GetAccountByBankIdAccountNumber(bankId, accountNumber int) (bankAccount.BankAccountInterface, error) {
	return user.customer.getAccountByNumber(accountNumber, bankId)
}
