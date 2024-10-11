package user

type Customer struct {
	totalBalance float64
}

func newCustomer(customerParameters map[string]interface{}) *Customer {
	return &Customer{
		totalBalance: 0,
	}
}
