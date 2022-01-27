// advanced-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	fmt.Println("In Logging:")
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			fmt.Println("In Logging: before the next middleware")
			f(w, r)
			fmt.Println("In Logging: after the next middleware")
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	fmt.Println("In Method:", m)
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				fmt.Println("In Method: NOT GET")
				return
			}

			// Call the next middleware/handler in chain
			fmt.Println("In Method: before the next middleware")
			f(w, r)
			fmt.Println("In Method: after the next middleware")
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		fmt.Println("In Chain:", m)
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Hello")

	fmt.Fprintln(w, "hello world")
}

func main() {
	fmt.Println("In main Hello:", Hello)
	fmt.Println("In main Logging:", Logging)
	fmt.Println("In main Method:", Method)
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
