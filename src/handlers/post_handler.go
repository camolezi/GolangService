package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/domain"
)

//PostHandler is the handler for /post url
type PostHandler struct {
	log     debug.Logger
	service servicesInterface
}

func (p *PostHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		p.getPost(writer, request)
	case http.MethodPost:
		p.addPost(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

//getPost is a function to handle GET requests at /post
func (p *PostHandler) getPost(writer http.ResponseWriter, request *http.Request) {

	//Get id from url- implement
	idString := strings.TrimPrefix(request.URL.Path, "/post/")

	if idString == "" {
		//Serve default page
		writer.Header().Set("Content-Type", "application/json")
		posts, _ := p.service.GetLatestPosts(10)

		postsJSON, err := json.Marshal(posts)

		if err != nil {
			p.log.Error().Println(err.Error())
		}

		writer.Write(postsJSON)
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)

	//Type error- id needs to be a number
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("id needs to be a number"))
		return
	}

	post, apiError := p.service.GetPost(id)
	//Api Error
	if apiError != nil {
		writer.WriteHeader(apiError.ErrorCode) //Write error header
		writer.Write([]byte(apiError.ErrorMessage))
		return
	}

	//for now trasnform to json here
	postJSON, errJSON := post.ToJSON()

	if errJSON != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		p.log.Error().Println(errJSON.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(postJSON)
}

func (p *PostHandler) addPost(writer http.ResponseWriter, request *http.Request) {
	//Get id from url- implement
	idString := strings.TrimPrefix(request.URL.Path, "/post")
	if idString != "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newPost := domain.Post{}

	//This need to be changed- must have a limit to size
	bodyData, _ := ioutil.ReadAll(request.Body)

	errJSON := newPost.FromJSON(bodyData)

	if errJSON != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//placeholder id
	postID := uint64(time.Now().Unix())
	newPost.ID = postID

	p.log.Debug().Printf("%v\n", newPost)

	//Try to create new post
	err := p.service.NewPost(postID, newPost)
	if err != nil {
		p.log.Error().Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)

}

//NewPostHandler return a new Post handler
func NewPostHandler(log debug.Logger) *PostHandler {
	return &PostHandler{log: log, service: &servicesWrapper{}}
}
