package main

import (
	"html/template"
	"log"
	"net/http"
)

type MoreInfo struct {
	Topic       string
	Description string
}

func moreInfo() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
	})

	http.HandleFunc("/moreinfo", func(w http.ResponseWriter, r *http.Request) {
		moreInfos := []MoreInfo{

			{Topic: "Basketball", Description: "Basketball is a team sport in which two teams, most commonly of five players each, opposing one another on a rectangular court, compete with the primary objective of shooting a basketball (approximately 9.4 inches (24 cm) in diameter) through the defender's hoop (a basket 18 inches (46 cm) in diameter mounted 10 feet (3.048 m) high to a backboard at each end of the court), while preventing the opposing team from shooting through their own hoop."},
		}
		moreInfoTemplate := `
			{{range .}}
				<li>{{.Topic}}: {{.Description}}</li>
			{{end}}

	`

		temp := template.Must(template.New("moreInfo").Parse(moreInfoTemplate))
		temp.Execute(w, moreInfos)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	moreInfo()
}
