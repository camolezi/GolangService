package model

import "encoding/json"

//User is a struct represinting a user - for now
type User struct {
	Login          string `json:"login"`
	Name           string `json:"title"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"password"` //this should not be in this struct in the future
}

//FromJSON Create a new user object from a json
func (p *User) FromJSON(data []byte) error {
	return json.Unmarshal(data, p)
}
