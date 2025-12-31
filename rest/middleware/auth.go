package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"

	"blog/util"
)

func (m *Middlewares) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		fmt.Println("Header Array: ", headerArr)

		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		bytArrSecret := []byte(m.cnf.JWTSecretKey)
		bytMessage := []byte(message)

		h := hmac.New(sha256.New, bytArrSecret)
		h.Write(bytMessage)

		hash := h.Sum(nil)
		newSignature := util.Base64UrlEncoder(hash)

		if newSignature != jwtSignature {
			http.Error(w, "Unauthorized. Tui Hacker", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
