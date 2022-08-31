package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var array []Customer

type CustomerHelper interface {
	Create()
	//Update()
}
type Customer struct {
	userid    string
	firstName string
	lastName  string
	birthDate time.Time
	gender    string
	email     string
	address   string
}

func Create(firstName string, lastName string, birthDate time.Time, gender string, email string, address string) (Customer, error) {
	customer := Customer{
		userid:    uuid.NewString(),
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		gender:    gender,
		email:     email,
		address:   address,
	}
	check := validation(customer)
	fmt.Println("Before Validation ", customer)
	if check == false {
		return customer, errors.New("error in validation")
	}
	array = append(array, customer)
	e := insert(customer)
	if e != nil {
		return customer, e
	}
	fmt.Println(array)
	return customer, nil
}

func Update(userid string, firstName string, lastName string, birthDate time.Time, gender string, email string, address string) (Customer, error) {
	customer := Customer{
		userid:    userid,
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		gender:    gender,
		email:     email,
		address:   address,
	}
	check := validation(customer)
	fmt.Println("Before Validation ", customer)
	if check == false {
		return customer, errors.New("Error in validation")
	}
	array = append(array, customer)
	e := update(customer)
	if e != nil {
		return customer, e
	}
	fmt.Println(array)
	return customer, nil
}
