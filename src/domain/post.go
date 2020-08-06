package domain

//Post is a struct representing one post
type Post struct {
	ID    uint64
	Title string
	Body  PostBody
}

//PostBody represents the body of a post, let this be just a string for now
type PostBody struct {
	text string
}
