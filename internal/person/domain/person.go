package domain

import "errors"

// ErrUnknown is used when a Family member not exist could not be found.
var ErrUnknown = errors.New("unknown Member")

type Relationship struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}

type Person struct {
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships"`
}

type Family struct {
	Members []Person `json:"members"`
}
