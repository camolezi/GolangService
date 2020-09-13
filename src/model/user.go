package model

import (
	"encoding/json"
	"time"
)

//User is a struct represinting a user - for now
type User struct {
	Login          string    `json:"login"`
	CreatedAt      time.Time `json:"createdAt"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"password"`
}

//FromJSON Create a new user object from a json
func (p *User) FromJSON(data []byte) error {
	return json.Unmarshal(data, p)
}
