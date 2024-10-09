package contact_manager

type user struct {
	userId    int
	firstName string
	lastName  string
	isAdmin   bool
	isActive  bool
	contacts  []contact
}

type contact struct {
	contactId int
	FirstName string
	LastName  string
	isActive  bool
	Details   []ContactDetails
}

type contactDetails struct {
	contactDetailsId int
	Type             string
	Email            string
	Number           int
}

///CRUD for Users

func NewUser() *user {

}

func GetAllUsers()
