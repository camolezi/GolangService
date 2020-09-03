package model

import (
	"encoding/json"
	"io"
)

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

//ToJSONData return JSON representation of the object
func (p *Post) ToJSONData() ([]byte, error) {
	return json.Marshal(*p)
}

//FromJSONData Create a new post object from a json
func (p *Post) FromJSONData(data []byte) error {
	return json.Unmarshal(data, p)
}

//FromIOReader Creates a new post from a IOReader
func (p *Post) FromIOReader(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	return decoder.Decode(p)
}
