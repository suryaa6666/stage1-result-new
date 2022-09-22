package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/add-blog.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

// var dataBlog = []
type Blog struct {
	Title     string
	Content   string
	Author    string
	Post_date string
}

var dataBlog = []Blog{}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Title : " + r.PostForm.Get("inputTitle")) // get value berdasarkan dari tag input name
	// fmt.Println("Content : " + r.PostForm.Get("inputContent"))

	var title = r.PostForm.Get("inputTitle")
	var content = r.PostForm.Get("inputContent")

	// let blog = {
	// 	title,
	// 	content
	// }

	var newBlog = Blog{
		Title:     title,
		Content:   content,
		Author:    "Samsul Rijal",
		Post_date: time.Now().String(),
	}

	// dataBlog.push(blog)
	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)
	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	response := map[string]interface{}{
		"Blogs": dataBlog,
	}

	tmpl.Execute(w, response)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	data := map[string]interface{}{
		"Title":   "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.",
		"Id":      id,
	}

	tmpl.Execute(w, data)
}
