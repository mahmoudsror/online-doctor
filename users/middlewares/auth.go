package middlewares

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/mahmoudsror/online-doctor/responses"
// 	"github.com/mahmoudsror/online-doctor/services/auth"
// )

// // IsAuthorized authorization middleware
// func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
// 	return func(response http.ResponseWriter, req *http.Request) {
// 		err := auth.ValidateToken(req)
// 		if err != nil {
// 			responses.Error(response, http.StatusUnauthorized, errors.New("unauthorized"))
// 			return
// 		}
// 		next(response, req)
// 	}
// }

// // JSONFormatter format response to json
// func JSONFormatter(next http.HandlerFunc) http.HandlerFunc {
// 	return func(response http.ResponseWriter, req *http.Request) {
// 		response.Header().Set("Content-Yype", "application/json")
// 		next(response, req)
// 	}
// }
