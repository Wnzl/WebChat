package auth

import (
	"github.com/Wnzl/webchat/models"
	"net/http"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var signingKey = []byte(os.Getenv("JWT_SECRET"))

// JwtMiddleware handler for jwt tokens
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// GetToken create a jwt token with user claims
func GetToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}

// GetUserClaimsFromContext return "user" claims as a map from request
func GetUserClaimsFromContext(req *http.Request) map[string]interface{} {
	//claims := context.Get(req, "user").(*jwt.Token).Claims.(jwt.MapClaims)
	claims := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims
}
