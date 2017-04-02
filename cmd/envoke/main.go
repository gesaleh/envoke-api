package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zbo14/envoke/api"
	// . "github.com/zbo14/envoke/common"
)

func main() {

	/*
		CreatePages(
			"create",
			"license_right",
			"login_register",
			"search",
			"verification",
		)

		RegisterTemplates(
			"create.html",
			"license_right.html",
			"login_register.html",
			"search.html",
			"verification.html",
		)

		// Create request multiplexer
		mux := http.NewServeMux()
		mux.HandleFunc("/create", TemplateHandler("create.html"))
		mux.HandleFunc("/license_right", TemplateHandler("license_right.html"))
		mux.HandleFunc("/login_register", TemplateHandler("login_register.html"))
		mux.HandleFunc("/verification", TemplateHandler("verification.html"))
		fs := http.Dir("static/")
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(fs)))
	*/

	// Create http router
	router := httprouter.New()

	// Create api and add routes to multiplexer
	api.NewApi().AddRoutes(router)

	// Start HTTP server with router
	http.ListenAndServe(":8888", router)

}
