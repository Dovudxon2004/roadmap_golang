package roadmap

import "strconv"

type ContactList struct {
	ID        int
	firstName string
	lastName  string
	phone     int
	email     string
	position  string
}

func NewContactList(ID int, first, last string, phone int, email, position string) *ContactList {
	return &ContactList{
		ID,
		first,
		last,
		phone,
		email,
		position,
	}
}

func (a *ContactList) Update(field, value string) {
	switch field {
	case "ID":
		a.ID, _ = strconv.Atoi(value)
	case "fristName":
		a.firstName = value
	case "lastName":
		a.lastName = value
	case "phone":
		a.phone, _ = strconv.Atoi(value)
	case "email":
		a.email = value
	case "position":
		a.position = value
	}
}

func (a *ContactList) Get(field string) any {
	switch field {
	case "ID":
		return a.ID
	case "fristName":
		return a.firstName
	case "lastName":
		return a.lastName
	case "phone":
		return a.phone
	case "email":
		return a.email
	case "position":
		return a.position
	default:
		return 0
	}
}

func (a *ContactList) GetAll() (int, string, string, int, string, string) {
	return a.ID, a.firstName, a.lastName, a.phone, a.email, a.position
}

func (a *ContactList) Delete(field string) {
	switch field {
	case "ID":
		a.ID = 0
	case "fristName":
		a.firstName = ""
	case "lastName":
		a.lastName = ""
	case "phone":
		a.phone = 0
	case "email":
		a.email = ""
	case "position":
		a.position = ""
	}
}

func StructInGolang() {

}
