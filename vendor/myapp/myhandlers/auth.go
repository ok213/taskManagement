package myhandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	tokenName = "mywebapp"
)

type keyContext string

// AuthHandler ...
func AuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// log.Println("Auth Handler Executing ...")
		// check if we have a cookie with out tokenName
		tokenCookie, err := r.Cookie(tokenName)
		if err != nil {
			log.Printf("Cookie parse error: %v\n", err)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		// just for the lulz, check if it is empty.. should fail on Parse anyway..
		if tokenCookie.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenCookie.Value, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return verifyKey, nil
		})

		if err != nil && !token.Valid {
			log.Printf("Token is not valid: %v \n", err)
			http.Redirect(w, r, "/signout", http.StatusSeeOther)
		}

		claims := token.Claims.(*MyCustomClaims)

		// !!! устанавливаем контекст !!!
		key := keyContext("Values")
		ctx := context.WithValue(r.Context(), key, claims)
		// log.Printf("---context send: %v\n", ctx)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)

}
