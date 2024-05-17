package middleware

import(
	"errors"
	"net/http"
	"github.com/Pranaenae/goauthapi/internal/tools"
	"github.com/Pranaenae/goauthapi/api"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		var username string = r.URL.query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error 
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
	}
}