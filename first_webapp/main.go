package main

import (
	"html/template"

	"net/http"
)

func process(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")

	t.Execute(writer, template.HTML(request.FormValue("comment"))) // 没有转义 的xxr攻击
	//t.Execute(writer, request.FormValue("comment"))                //有转义 无xxr攻击
}
func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()

}
