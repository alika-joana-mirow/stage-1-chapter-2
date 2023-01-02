package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title": "Personal Web",
}

type project struct {
	Id        int
	Name      string
	StartDate string
	EndDate   string
	Duration  string
	Tech      []string
	Desc      string
}

var projects = []project{}

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/home", home).Methods("GET").Name("home")
	route.HandleFunc("/add-project", formProject).Methods("GET")
	route.HandleFunc("/home", addProject).Methods("POST")
	route.HandleFunc("/detail-project/{id}", detailProject).Methods("GET")
	route.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET")
	route.HandleFunc("/edit-project/{id}", editProject).Methods("GET")
	route.HandleFunc("/edit-project-form/{id}", editProjectForm).Methods("POST")
	route.HandleFunc("/contact-me", contactMe).Methods("GET")

	fmt.Println("Server is running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	respData := map[string]interface{}{
		"Data":     Data,
		"projects": projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/add-my-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Project Name : " + r.PostForm.Get("name"))
	// fmt.Println("Start Date : " + r.PostForm.Get("start-date"))
	// fmt.Println("End Date : " + r.PostForm.Get("end-date"))
	// fmt.Println("Description : " + r.PostForm.Get("desc"))
	// fmt.Println("Technologies : ", r.Form["tech"])

	name := r.PostForm.Get("name")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	desc := r.PostForm.Get("desc")
	tech := r.Form["tech"]
	duration := proDuration(startDate, endDate)

	var newProject = project{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Desc:      desc,
		Tech:      tech,
		Duration:  duration,
	}

	projects = append(projects, newProject)

	// fmt.Println(projects)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// fmt.Println(id)

	projects = append(projects[:id], projects[id+1:]...)

	http.Redirect(w, r, "/home", http.StatusFound)
}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	tmpl, err := template.ParseFiles("views/detail-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	respDetail := map[string]interface{}{
		"Data": Data,
		"id":   id,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respDetail)
}

func editProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/edit-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	proDetail := project{}

	for index, data := range projects {
		if index == id {
			proDetail = project{
				Id:        id,
				Name:      data.Name,
				StartDate: data.StartDate,
				EndDate:   data.EndDate,
				Desc:      data.Desc,
				Tech:      data.Tech,
				Duration:  data.Duration,
			}
		}
	}

	respData := map[string]interface{}{
		"Data":      Data,
		"proDetail": proDetail,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func editProjectForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := r.PostForm.Get("name")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	desc := r.PostForm.Get("desc")
	tech := r.Form["tech"]
	duration := proDuration(startDate, endDate)

	var newProject = project{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Desc:      desc,
		Tech:      tech,
		Duration:  duration,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	projects[id] = newProject

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/contact-me.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func proDuration(StartDate, EndDate string) string {
	sDate, _ := time.Parse("2006-01-02", StartDate)

	eDate, _ := time.Parse("2006-01-02", EndDate)

	// substracting date or time
	interval := eDate.Sub(sDate)

	year := int(interval.Hours() / (12 * 30 * 24))
	month := int(interval.Hours() / (20 * 24))
	week := int(interval.Hours() / (7 * 24))
	day := int(interval.Hours() / 24)

	if year != 0 {
		return "Duration : " + strconv.Itoa(year) + "Year"
	}
	if month != 0 {
		return "Duration : " + strconv.Itoa(month) + "Month"
	}
	if week != 0 {
		return "Duration : " + strconv.Itoa(week) + "Week"
	} else {
		return "Duration : " + strconv.Itoa(day) + "Day"
	}
}
