package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers/response"
	"github.com/camolezi/MicroservicesGolang/src/model"
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
func (p *PostHandler) getPost(defaultWriter http.ResponseWriter, request *http.Request) {
	response := response.CreateResponse(defaultWriter, p.log)

	//Get id from url- implement
	idString := strings.TrimPrefix(request.URL.Path, "/post/")

	if idString == "" {
		//Serve default page
		posts, _ := p.service.GetLatestPosts(10)
		postsJSON, err := json.Marshal(posts)

		if err != nil {
			p.log.Error().Println(err.Error())
		}

		response.WriteJSON(postsJSON)
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)

	//Type error- id needs to be a number
	if err != nil {
		response.BadRequest("id needs to be a number")
		return
	}

	post, apiError := p.service.GetPost(int64(id))
	//Api Error
	if apiError != nil {
		response.WriteError(apiError.ErrorCode, apiError.ErrorMessage)
		return
	}

	//for now trasnform to json here
	postJSON, errJSON := post.ToJSONData()

	if errJSON != nil {
		response.WriteError(http.StatusInternalServerError, errJSON.Error())
		p.log.Error().Println(errJSON.Error())
		return
	}

	response.WriteJSON(postJSON)
}

func (p *PostHandler) addPost(defaultWriter http.ResponseWriter, request *http.Request) {
	response := response.CreateResponse(defaultWriter, p.log)

	//Get id from url- implement
	idString := strings.TrimPrefix(request.URL.Path, "/post")
	if idString != "" {
		response.WriteError(http.StatusBadRequest, "")
		return
	}

	newPost := model.Post{}
	errJSON := newPost.FromIOReader(request.Body)

	if errJSON != nil {
		response.BadRequest(errJSON.Error())
		return
	}

	//placeholder id
	postID := int64(time.Now().Unix()) //This should be probably be coming from the database
	newPost.ID = postID

	//Try to create new post
	err := p.service.NewPost(postID, newPost)
	if err != nil {
		response.ServerError(err.Error())
		return
	}

	//Return the newly craeted object
	returnData, err := newPost.ToJSONData()
	if err != nil {
		response.ServerError(err.Error())
		return
	}

	response.Created(returnData)
}

//NewPostHandler return a new Post handler
func NewPostHandler(log debug.Logger) *PostHandler {
	return &PostHandler{log: log, service: &servicesWrapper{}}
}
