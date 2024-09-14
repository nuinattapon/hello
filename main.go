package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Hello, 世界, สวัสดี")
	fmt.Fprintln(w, "Version 2.4")

	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "Server Name: %s\n", hostName)
	addrs, _ := net.LookupIP(hostName)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Fprintf(w, "Server Addr: %s\n", ipv4)
		}
	}
	fmt.Fprintln(w, "Remote Addr:", r.RemoteAddr)
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URI:", r.URL.RequestURI())
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		startTime.Year(), startTime.Month(), startTime.Day(),
		startTime.Hour(), startTime.Minute(), startTime.Second())
	fmt.Fprintln(w, "Date:", formatted)

	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	user := User{
		Id:    1,
		Name:  "Somsri Phrapradaeng",
		Email: "somsri.phrapradaeng@fakemail.com",
		Phone: "0812374651",
	}
	json.NewEncoder(w).Encode(user)
	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	temp, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}
	user := User{
		Id:    1,
		Name:  "Somsri Phrapradaeng",
		Email: "somsri.phrapradaeng@fakemail.com",
		Phone: "0812374651",
	}
	temp.Execute(w, user)

	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "pong")
	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0

	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "2.4")
	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0

	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/json", jsonHandler)
	mux.HandleFunc("/template", templateHandler)
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/version", versionHandler)
	mux.HandleFunc("/fibo", fiboHandler)
	mux.HandleFunc("/fibo/", fiboHandler)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}

}

func fiboHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	nStr := strings.TrimPrefix(r.URL.Path, "/fibo/")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Fprintf(w, "Please provide a valid number. Fibonacci calculation request is rejected!\n")
	} else {
		if n <= 45 && n >= 0 {
			fib := fibo(n)
			fmt.Fprintf(w, "Fibonacci(%d)=%d\n", n, fib)
		} else {
			fmt.Fprintf(w, "%d is not valid. Please provide number <= 45, >= 0. Fibonacci calculation request is rejected!\n", n)
		}
	}
	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0

	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
