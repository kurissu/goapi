package middleware

import (
	"error"
	"net/http"

	"github.com/kurissu/goapi"
	"github.com/kurissu/goapi/api"
	"github.com/kurissu/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)


var UnAuthorizationError = errors.New("Invalid username or token.")

func Authrization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error 

		if username == "" || token == "" {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
			return
		}

		var database *tools.DatabaseInterface
		database, err =tools.NewDatabase()
		if err != nil {
			api.internalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if(loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
		}

		next.ServeHTTP(w, r)

	})
}		