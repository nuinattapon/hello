package handlers

import (
	"encoding/json"
	"fmt"
	"hello/config"
	"html/template"
	"net"
	"net/http"
	"os"
	"time"
)

// User represents a user entity
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// HomeHandler handles requests to the root path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Hello, 世界, สวัสดี")
	fmt.Fprintln(w, "Version", config.Version)

	hostName, err := os.Hostname()
	if err != nil {
		// Log the error but continue
		fmt.Printf("Error getting hostname: %v\n", err)
		hostName = "unknown"
	}
	fmt.Fprintf(w, "Server Name: %s\n", hostName)

	addrs, err := net.LookupIP(hostName)
	if err != nil {
		// Log the error but continue
		fmt.Printf("Error looking up IP: %v\n", err)
	} else {
		for _, addr := range addrs {
			if ipv4 := addr.To4(); ipv4 != nil {
				fmt.Fprintf(w, "Server Addr: %s\n", ipv4)
			}
		}
	}

	fmt.Fprintln(w, "Remote Addr:", r.RemoteAddr)
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URI:", r.URL.RequestURI())

	now := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Fprintln(w, "Date:", formatted)
}

// JSONHandler returns user data in JSON format
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	user := User{
		Id:    1,
		Name:  "Sam Phrapradaeng",
		Email: "sam.phrapradaeng@fakemail.com",
		Phone: "+66-81-234-5678",
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		// Log the error
		fmt.Printf("Error encoding JSON: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// TemplateHandler renders an HTML template with user data
func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	temp, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	user := User{
		Id:    1,
		Name:  "Sam Phrapradaeng",
		Email: "sam.phrapradaeng@fakemail.com",
		Phone: "+66-81-234-5678",
	}

	if err := temp.Execute(w, user); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// PingHandler responds with "pong"
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "pong")
}

// VersionHandler returns the application version
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, config.Version)
}
