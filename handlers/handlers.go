package handlers

import (
	"github.com/course_spec/data_control/autorization/cookies"
	"html/template"
	"net/http"
)




var LoginPageHandler = http.HandlerFunc(
	func(w http.ResponseWriter,r *http.Request){
	parsedTemplate,_ := template.ParseFiles("templates/login.html")
	parsedTemplate.Execute(w,nil)
	})

var HomePageHandler = http.HandlerFunc(
	func(w http.ResponseWriter,r *http.Request){
	UserName := cookies.GetUserName(r)
	if UserName != ""{
		data := map[string]interface{}{
			//tyt tozhe nyzhno s bol'shoi
			"UserName":UserName,
		}
		parsedTemplate,_ := template.ParseFiles("templates/home.html")
		parsedTemplate.Execute(w,data)
	} else {
		http.Redirect(w,r,"/",302)
	}

})

var LoginHomePageHandler = http.HandlerFunc(
 	func(w http.ResponseWriter,r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	target   := "/"
	if username != "" && password != ""{
		cookies.SetsSession(username,w)
		target = "/home"
	}
	http.Redirect(w,r,target,302)
})

var LogoutFormPageHandler = func(w http.ResponseWriter,r *http.Request){
	cookies.ClearSession(w)
	http.Redirect(w,r,"/",302)
}
