package main

import (
	"encoding/json"
	"hello-world/config"
	"html/template"
	"log"
	"net/http"
)

// config
var cfg *config.Config

func setupServer() *http.Server {
	// Load configuration
	cfg = config.LoadConfig()

	// Create new mux
	mux := http.NewServeMux()

	// Handle static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/about", handleAbout)
	mux.HandleFunc("/ping", ping)

	// Create server with custom 404 handler
	server := &http.Server{
		Addr: ":" + cfg.AppPort,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the path exists in mux
			h, pattern := mux.Handler(r)
			if pattern == "/" && r.URL.Path != "/" {
				logIncomingRequest(r)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404"))
				return
			}
			h.ServeHTTP(w, r)
		}),
	}

	return server
}

func main() {
	server := setupServer()

	log.Printf("Starting %s version %s on http://localhost:%s\n",
		cfg.AppName, cfg.AppVersion, cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	logIncomingRequest(r)

	type PageData struct {
		Title       string
		Message     string
		AppName     string
		AppVersion  string
		Environment string
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := PageData{
		Title:       "Welcome",
		Message:     "Hello, World!",
		AppName:     cfg.AppName,
		AppVersion:  cfg.AppVersion,
		Environment: cfg.AppEnv,
	}

	tmpl.Execute(w, data)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	logIncomingRequest(r)

	type PageData struct {
		Title   string
		Message string
	}

	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	data := PageData{
		Title:   "About Us",
		Message: "This is a simple Go web application demonstrating routing and templates.",
	}

	tmpl.Execute(w, data)
}

func ping(w http.ResponseWriter, r *http.Request) {
	logIncomingRequest(r)

	type Response struct {
		Message   string `json:"message"`
		Ip        string `json:"ip"`
		UserAgent string `json:"useragent"`
	}

	response := Response{
		Message:   "pong",
		Ip:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
	}

	json.NewEncoder(w).Encode(response)
}

func logIncomingRequest(r *http.Request) {
	log.Printf("%s %s %s %s\n", r.UserAgent(), r.RemoteAddr, r.Method, r.URL)
}
