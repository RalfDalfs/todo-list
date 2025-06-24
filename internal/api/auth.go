package api

import (
	"net/http"

	"github.com/RalfDalfs/todo-list/internal/config"
)

func auth(next http.HandlerFunc, cfg *config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		pass := cfg.Passw
		if len(pass) > 0 {
			var jwt string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}
			var valid bool
			// здесь код для валидации и проверки JWT-токена
			tokenTrue, err := jwtGen(pass)
			if err != nil {
				sendJSONError(w, err.Error(), http.StatusInternalServerError)
			}
			if jwt == tokenTrue {
				valid = true
			}

			if !valid {
				// возвращаем ошибку авторизации 401
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
