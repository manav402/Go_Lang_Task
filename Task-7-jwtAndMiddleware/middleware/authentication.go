package middleware

import (
	"manav402/crudBooks/database"
	"net/http"

	"github.com/golang-jwt/jwt"
)

// generator function generate hash for key with all the claims
func GenerateJwt(role string) (string, error) {
	// secret key from env file
	var secretKey, err = database.GetEnv("JWTSECRET")
	if err != nil {
		return "", err
	}

	// creating a claim with role which will needed for authentication
	var claims = &jwt.MapClaims{
		"data": map[string]string{
			"role": string(role),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// verify jwt is a middleware which is used for checking if jwt is valid or not
func VerifyJwt(next http.HandlerFunc, allwedRole string) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		cookieMap, err := req.Cookie("cookie")
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("your session is expired please login err:=" + err.Error()))
			return
		}

		jwtStr := cookieMap.Value
		token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {

			var secretKey, err = database.GetEnv("JWTSECRET")
			if err != nil {
				return nil, err
			}

			return []byte(secretKey), nil
		})

		// if err means cant parse the jwt
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("can't parse the jwt err :"+err.Error()))
			return
		}

		// if token is invalis formate
		if !token.Valid {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("token is not valid"))
		} else {
			// getting claims
			claims := token.Claims.(jwt.MapClaims)
			data := claims["data"].(map[string]interface{})
			role := data["role"].(string)

			if allwedRole == "any" {
				next(res, req)
			} else if role == allwedRole {
				next(res, req)
			} else {
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte("you are unauthorized to access this site"))
				return
			}
		}
	})
}
