package main

import (
	"net/mail"
	"time"
)

func getAge(Bday time.Time) int {
	curTime := time.Now()
	cty, ctm, ctd := curTime.Date()
	bdy, bdm, bdd := Bday.Date()
	age := cty - bdy
	agemonth := ctm - bdm
	agedate := ctd - bdd
	if agemonth < 0 {
		age--
	}
	if agemonth == 0 && agedate < 0 {
		age--
	}
	return age
}

func validation(customer Customer) bool {
	if len(customer.firstName) > 100 || len(customer.firstName) == 0 {
		return false
	}
	if len(customer.lastName) > 100 || len(customer.lastName) == 0 {
		return false
	}
	if customer.gender != "Male" && customer.gender != "Female" || len(customer.gender) == 0 {
		return false
	}
	_, err := mail.ParseAddress(customer.email)
	if err != nil {
		return false
	}
	if len(customer.address) > 200 {
		return false
	}
	age := getAge(customer.birthDate)
	if age < 18 || age > 60 {
		return false
	}
	return true
}
