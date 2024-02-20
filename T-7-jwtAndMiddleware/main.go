// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// )

// var secret = []byte("mySecret")

// func LoginHandler(res http.ResponseWriter, req *http.Request) {

// 	unsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
// 		Issuer:    "tridip",
// 		ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
// 	})
// 	signed, err := unsigned.SignedString(secret)

// 	if err != nil && err == jwt.ErrInvalidKey {
// 		res.Write([]byte("Token Error"))
// 		return
// 	}
// 	http.SetCookie(res,
// 		&http.Cookie{
// 			Name:  "Tridip",
// 			Value: signed,
// 		},
// 	)

// }
// func ValidateJwt(next http.HandlerFunc) http.Handler {
// 	fmt.Println("hello")
// 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		val, err := req.Cookie("Tridip")
// 		if err != nil {
// 			fmt.Println("no token found here ")
// 			return
// 		}
// 		tkn,err := jwt.ParseWithClaims(val.Value,&jwt.StandardClaims{},func (Token *jwt.Token)(interface{},error){
	
// 			return secret,nil
// 		})

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if !tkn.Valid{
// 			res.Write([]byte("why the fuck are you here go away"))
// 			return
// 		}
// 		next(res,req)
// 	})
// }

// func homeHandler(res http.ResponseWriter, req *http.Request) {
// 	res.Write([]byte("hello and welcome to brawl talk"))
// }
// func main() {
// 	http.Handle("/home", ValidateJwt(homeHandler))
// 	http.HandleFunc("/login", LoginHandler)
// 	http.ListenAndServe(":8000", nil)
// }

// // // type myInter interface{
// // // 	hello()
// // // }

// // // type myStr struct{
// // // 	myInter myInter
// // // }

// // var secret = []byte("manav")

// // type Message struct {
// // 	Status string `json:"status"`
// // 	Info   string `json:"info"`
// // }

// // func GenerateJWT() (string, error) {
// // 	// myStr.myInter.hello()
// // 	msg := jwt.New(jwt.SigningMethodHS256)
// // 	// var a interface{}=[]byte("aSecretKey")
// // 	fmt.Println(secret)
// // 	str, err := msg.SignedString(secret)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return "", err
// // 	}
// // 	return str, nil
// // }

// // func ValidateJwt(next http.HandlerFunc) http.Handler {
// // 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// // 		if req.Header["Token"] != nil {
// // 			// x := &Message{} // *Message
// // 			token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// // 				// type assertion is happening here
// // 				_, ok := token.Method.(*jwt.SigningMethodHMAC)
// // 				if !ok {
// // 					_, err := res.Write([]byte("aww you are not authenticated!!!"))
// // 					if err != nil {
// // 						res.Write([]byte("error parsing token"))
// // 						return nil, nil
// // 					}
// // 				}
// // 				return nil, nil
// // 			})
// // 			if err != nil {
// // 				fmt.Println(err)
// // 			}
// // 			fmt.Println(token)
// // 			if token.Valid {
// // 				next(res, req)
// // 			} else {
// // 				res.Write([]byte("your token is invalid"))
// // 				return
// // 			}
// // 			fmt.Println(token, err)
// // 			return
// // 		} else {
// // 			res.Write([]byte("no token found in header"))
// // 			return
// // 		}
// // 		// next(res, req)
// // 	})
// // }

// // // func middleware2(next http.HandlerFunc) http.HandlerFunc{
// // // 	return http.HandlerFunc(func(res http.ResponseWriter,req *http.Request){
// // // 		fmt.Println("hello i am middleware 2")
// // // 		next(res,req)
// // // 	})
// // // }

// // // func middleware1(next http.HandlerFunc) http.Handler {
// // // 	return http.HandlerFunc(func(res http.ResponseWriter,req *http.Request){
// // // 		fmt.Println("i am a middleware1")
// // // 		// middleware2(next)
// // // 		next(res,req)
// // // 	})
// // // }

// // func AuthPage(res http.ResponseWriter, req *http.Request) {
// // 	fmt.Println("hello")
// // 	token, err := GenerateJWT()
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	// client := &http.Client{}
// // 	req.Header.Set("Token",token)
// // 	fmt.Println(token)
// // 	res.Write([]byte("heyy you are now have a valid token for login "+token))
// // }

// // func helloHanlder(res http.ResponseWriter, req *http.Request) { // newResponse := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// // 	var message Message
// // 	fmt.Println("got req ", req.Method, req.URL)
// // 	res.Header().Set("Content-Type", "application/json")
// // 	json.NewDecoder(req.Body).Decode(&message)
// // 	json.NewEncoder(res).Encode(&message)
// // }

// // func main() {
// // 	http.Handle("/hello", ValidateJwt(helloHanlder))
// // 	http.HandleFunc("/login", AuthPage)
// // 	http.ListenAndServe(":8000", nil)
// // }
