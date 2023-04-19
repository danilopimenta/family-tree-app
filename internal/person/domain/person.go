package domain

import "errors"

// PersonNotFound is used when a Family member not exist could not be found.
var PersonNotFound = errors.New("person not found")

// Relationship is the relationship between two people.
// the only possibilities are "parent" for now.
type Relationship struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}

// Person is a member of a family.
type Person struct {
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships"`
}

// Family a list of Members.
type Family struct {
	Members []Person `json:"members"`
}
