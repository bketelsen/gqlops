package model 

import "strings"

Profile: {
	ID:  <=10
	Name: strings.ContainsAny("z")
} @go(,complete=Complete)
