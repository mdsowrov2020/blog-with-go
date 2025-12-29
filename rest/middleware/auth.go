package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"strings"

	"blog/config"
	"blog/util"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header().Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ")

		if len(header) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		accessTokenArr := strings.Split(accessToken, ".")

		if len(accessTokenArr) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenHeader := accessTokenArr[0]
		tokenPayload := accessTokenArr[1]
		tokenSignature := accessTokenArr[2]

		message := tokenHeader + "." + tokenPayload
		cnf := config.GetConfig()
		bytArrSecret := []byte(cnf.JWTSecretKey)
		bytMessage := []byte(message)

		h := hmac.New(sha256.New, bytArrSecret)
		h.Write(bytMessage)

		hash := h.Sum(nil)
		newSignature := util.Base64UrlEncoder(hash)

		if newSignature != tokenSignature {
			http.Error(w, "Unauthorized. Tui Hacker", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
