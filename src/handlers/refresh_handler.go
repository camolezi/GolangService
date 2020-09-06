package handlers

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
)

//RefreshHandler is the handler to refresh access token
type RefreshHandler struct {
	Log debug.Logger
}

func (u *RefreshHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	u.refreshToken(writer, request)
}

func (u *RefreshHandler) refreshToken(defaultWriter http.ResponseWriter, request *http.Request) {

}
