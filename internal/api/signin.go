package api

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"

	"github.com/RalfDalfs/todo-list/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func signinHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendJSONError(w, "No allowed method", http.StatusMethodNotAllowed)
			return
		}

		type pass struct {
			Password string `json:"password"`
		}
		var p pass

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		pAcc := cfg.Passw
		signedToken, err := jwtGen(pAcc)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusInternalServerError)
		}
		if pAcc == p.Password {
			writeJSON(w, map[string]interface{}{"token": signedToken}, http.StatusOK)
			return
		} else {
			sendJSONError(w, "Invalid password", http.StatusUnauthorized)
			return
		}
	}
}

func jwtGen(pAcc string) (string, error) {
	secret := []byte(pAcc)
	hash := sha256.Sum256(secret)

	claims := jwt.MapClaims{
		"hash": hash,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
