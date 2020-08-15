package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/domain"
	servicesPkg "github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

type servicesInterface interface {
	GetPost(id uint64) (domain.Post, *utils.ErrorAPI)
}

var (
	service servicesInterface
)

type services struct{}

func (s *services) GetPost(id uint64) (domain.Post, *utils.ErrorAPI) {
	return servicesPkg.GetPost(id)
}

func init() {
	service = &services{}
}

//PostHandler is the handler for /post url
type PostHandler struct{}

func (p *PostHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getPost(writer, request)
	case http.MethodPost:
		addPost(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

//getPost is a function to handle GET requests at /post
func getPost(writer http.ResponseWriter, request *http.Request) {

	//Get id from url- implement
	idString := strings.TrimPrefix(request.URL.Path, "/post/")

	if idString == "" {
		//Serve default page
		writer.Write([]byte("Default post will be here, or all posts"))
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)

	//Type error- id needs to be a number
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("id needs to be a number"))
		return
	}

	post, apiError := service.GetPost(id)
	//Api Error
	if apiError != nil {
		writer.WriteHeader(apiError.ErrorCode) //Write error header
		writer.Write([]byte(apiError.ErrorMessage))
		return
	}

	//for now trasnform to json here
	postJSON, errJSON := json.Marshal(post)

	if errJSON != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(postJSON)
}

func addPost(writer http.ResponseWriter, request *http.Request) {
	//Get id from url- implement
	idString := strings.TrimPrefix(request.URL.Path, "/post/")
	if idString != "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//placeholder id
	postID := time.Now().Unix()

	newPost := domain.Post{}

	bodyData, _ := ioutil.ReadAll(request.Body)

	errJSON := json.Unmarshal(bodyData, &newPost)

	if errJSON != nil {
		log.Println("Error in post: " + errJSON.Error())
		log.Println("data recived: " + string(bodyData))
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		log.Println(newPost.Title)
		log.Println(postID)
	}

}

//NewPostHandler return a new Post handler
func NewPostHandler() *PostHandler {
	return &PostHandler{}
}
