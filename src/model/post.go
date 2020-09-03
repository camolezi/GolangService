package model

import "encoding/json"

//Post is a struct representing one post
type Post struct {
	ID    uint64   `json:"id"`
	Title string   `json:"title"`
	Body  PostBody `json:"body"`
}

//PostBody represents the body of a post, let this be just a string for now
type PostBody struct {
	Text string `json:"text"`
}

//ToJSON return JSON representation of the object
func (p *Post) ToJSON() ([]byte, error) {
	return json.Marshal(*p)
}

//FromJSON Create a new post object from a json
func (p *Post) FromJSON(data []byte) error {
	return json.Unmarshal(data, p)
}
