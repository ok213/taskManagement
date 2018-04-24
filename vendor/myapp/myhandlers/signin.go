package myhandlers

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// location of the files used for signing and verification
const (
	privKeyPath = "vendor/myapp/keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "vendor/myapp/keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// keys are held in global variables
// i havn't seen a memory corruption/info leakage in go yet
// but maybe it's a better idea, just to store the public key in ram?
// and load the signKey on every signing request? depends on  your usage i guess
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// MyCustomClaims ...
type MyCustomClaims struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	ShortName  string `json:"shortName"`
	Department string `json:"department"`
	jwt.StandardClaims
}

type pageSigninValues struct {
	MyCustomClaims
	Error string
}

// Signin ...
func (c *AppContext) Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("Signin Handler Executing ...")
	claims := MyCustomClaims{}
	pageVals := pageSigninValues{claims, ""}
	switch r.Method {
	case "GET":
		// отдать форму для аворизации
		c.Tmpl.ExecuteTemplate(w, "signin", pageVals)
	case "POST":
		// значит пришли какие-то данные, нужно их проверить
		r.ParseForm()
		user, err := c.Db.GetUser(r.Form.Get("email"))
		// если любая ошибка, то записать ее в лог и выдать сообщение о некой ошибке
		if err != nil {
			log.Printf("Signin Handler Error: %v\n", err.Error())
			pageVals.Error = "Error: "
			c.Tmpl.ExecuteTemplate(w, "signin", pageVals)
			return
		}
		if user.Hash != r.Form.Get("password") {
			log.Printf("Signin Handler Error: Incorrect password for the user '%s'\n", user.Email)
			pageVals.Error = "Error: "
			c.Tmpl.ExecuteTemplate(w, "signin", pageVals)
			return
		}

		// логин и пароль прошли проверку, след. создаем токен доступа
		claims = MyCustomClaims{
			user.ID,
			user.Email,
			user.Role,
			user.FirstName,
			user.LastName,
			user.MiddleName,
			user.ShortName,
			user.Department,
			jwt.StandardClaims{
				Id:        randStringRunes(32),
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 9).Unix(),
				Issuer:    "mywebapp",
			},
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("RS512"), claims)
		tokenString, err := token.SignedString(signKey)
		if err != nil {
			log.Printf("Token Signing error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			pageVals.Error = "Error: while Signing Token!"
			c.Tmpl.ExecuteTemplate(w, "signin", pageVals)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  tokenName,
			Value: tokenString,
			Path:  "/",
			// MaxAge: time.Now().Add(time.Hour * 9).Unix(),
			MaxAge:   32400,
			HttpOnly: true,
		})

		// http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Redirect(w, r, "/incoming", http.StatusSeeOther)

	default:
		http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusBadRequest)
	}

}
