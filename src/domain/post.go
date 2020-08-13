package domain

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
