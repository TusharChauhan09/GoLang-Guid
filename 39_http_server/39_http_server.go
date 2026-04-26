// 39_http_server.go
// Topic: net/http server — handlers, mux, middleware
//
// SIMPLEST SERVER
//   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//       fmt.Fprintln(w, "hello")
//   })
//   log.Fatal(http.ListenAndServe(":8080", nil))
//
// HANDLER INTERFACE
//   type Handler interface {
//       ServeHTTP(ResponseWriter, *Request)
//   }
//   HandlerFunc — func(w, r) adapter so plain funcs satisfy Handler.
//
// MUX (router)
//   mux := http.NewServeMux()
//   mux.HandleFunc("/path", fn)
//   Go 1.22+ supports method+pattern: mux.HandleFunc("GET /users/{id}", fn)
//   r.PathValue("id") fetches the path variable.
//
// REQUEST DATA
//   r.Method, r.URL.Path, r.URL.Query().Get("k")
//   r.Header.Get("X-Foo")
//   r.ParseForm(); r.FormValue("k")
//   json.NewDecoder(r.Body).Decode(&v)
//
// RESPONSE
//   w.Header().Set("Content-Type", "application/json")
//   w.WriteHeader(http.StatusCreated)       // call BEFORE Write
//   w.Write([]byte("..."))   /  fmt.Fprint(w, ...)
//
// SERVER WITH OPTIONS
//   srv := &http.Server{
//       Addr:         ":8080",
//       Handler:      mux,
//       ReadTimeout:  5 * time.Second,
//       WriteTimeout: 10 * time.Second,
//   }
//   srv.ListenAndServe()
//
// MIDDLEWARE PATTERN
//   func logger(next http.Handler) http.Handler {
//       return http.HandlerFunc(func(w, r) {
//           log.Println(r.Method, r.URL.Path)
//           next.ServeHTTP(w, r)
//       })
//   }
//
// To run: change main2 -> main, then `go run 39_http_server.go`
// Curl: curl localhost:8080/hello?name=Ada

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type apiResp struct {
	Message string `json:"message"`
	When    string `json:"when"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResp{
		Message: "hello " + name,
		When:    time.Now().Format(time.RFC3339),
	})
}

// Middleware
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Rename main -> main2 to avoid conflict with other files when compiled together.
// To actually run server, rename this back to main and remove others or build it alone:
//    go run 39_http_server.go
func main() {
	mux := http.NewServeMux()

	// Go 1.22+ method routing
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "user id: %s\n", id)
	})
	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      logger(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
