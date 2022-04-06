package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"jobapp.com/m/common"
)

var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

func UpdateContextUserModel(c *gin.Context, userModel UserModel) error {
	c.Set("my_user_model", userModel)
	return nil

}

func AuthMiddleware(useMiddleware bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !useMiddleware {
			c.Next()
			return
		}
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.SecretKeyJwt))
			return b, nil
		})
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := (claims["user_name"].(string))
			model, err := FindOneSql(username)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			UpdateContextUserModel(c, model)
		}
	}
}

func GenToken(username string) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"user_name": username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(common.SecretKeyJwt))
	return token
}
