package middlewares

import (
	"errors"
	"go-authentication/dtos"
	"go-authentication/infrastructure"
	"go-authentication/utils"
	"net/http"
)

func HTTPanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logger := infrastructure.NewLogger()
				logger.Err(r.Context(), "panic occured", errors.New("panic occured"), rvr)

				utils.WriteResponse(w, dtos.Response{
					StatusCode: http.StatusInternalServerError,
					Error:      "Internal Server Error occured",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
