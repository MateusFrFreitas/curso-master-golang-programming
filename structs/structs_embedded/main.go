package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "Jogn Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Streen 20, London",
			phone:   3490345938439048,
		},
	}

	fmt.Printf("%+v\n", john)

	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "test@test.com", phone: 439058390458, address: "Sidney, Australia"}
	fmt.Println(myContact)
}
