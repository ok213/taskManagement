package myhandlers

import (
	"log"
	"net/http"
)

// Signout ...
func (c *AppContext) Signout(w http.ResponseWriter, r *http.Request) {
	log.Println("Signout HandlerFunc executing...")
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", 302)
	return
}
