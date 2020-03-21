package middlewares

import (
	"errors"
	"net/http"

	"github.com/mahmoudsror/online-doctor/users/responses"
	"github.com/mahmoudsror/online-doctor/users/services/auth"
)

// IsAuthorized authorization middleware
func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		err := auth.ValidateToken(req)
		if err != nil {
			responses.Error(response, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(response, req)
	}
}
