package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
)
//opredelenie cookov standar razmeri 64/32
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetUserName(request *http.Request)(userName string){
	cookie,err := request.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session",cookie.Value,&cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

func SetsSession(userName string,response http.ResponseWriter){
	value := map[string]string{"username":userName}
	encoded, err := cookieHandler.Encode("session",value)
	if err == nil {
		cookie := &http.Cookie{
			Name: "session",
			Value: encoded,
			Path: "/",
		}
		http.SetCookie(response,cookie)
	}
}

func ClearSession(response http.ResponseWriter){
	cookie := &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(response,cookie)
}

