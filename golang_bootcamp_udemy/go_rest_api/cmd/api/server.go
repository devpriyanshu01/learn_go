package main

import (
	"crypto/tls"
	"encoding/json"
	"strconv"
	"strings"

	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"time"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	port := ":3001"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teacherHandler)

	mux.HandleFunc("/students/", studentHandler)

	mux.HandleFunc("/execs/", execHandler)

	//Configuring TLS handshake begins.
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Second*30)

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	// secureMux := ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)

	//create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler: secureMux,
		// Handler:   rl.Middleware(mw.SecurityHeaders(mux)),
		Handler: mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux)))))),

		// Handler:   mw.Hpp(hppOptions)(rl.Middleware(mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mux))))),
		// Handler:   middlewares.Cors(mux),
		// Handler: middlewares.ResponseTimeMiddleware(mux),
		// Handler: middlewares.Compression(mux),
		TLSConfig: tlsConfig,
	}

	// Middleware is a function that wraps an http.Handler with additional functionality
	// type Middleware func(http.Handler) http.Handler

	// func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) (http.Handler) {
	// 	for _, middleware := range middlewares {
	// 		handler = middleware(handler)
	// 	}
	// 	return handler
	// }

	fmt.Println("server is running on port:", port)
	// err := http.ListenAndServe(port, nil)
	err := server.ListenAndServeTLS(cert, key)
	//configuring tls handshake ends.
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root route"))
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)

	case http.MethodPost:
		fmt.Fprintf(w, "post method /teachers")
	}
}

// Creating a Teacher Struct for in-memory Teacher database.
type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Class     string
	Subject   string
}

var teachers = make(map[int]Teacher)

// var mutex = &sync.Mutex{}
var nextID = 1

// initialize some dummy teachers data.
func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9F",
		Subject:   "Maths",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Karl",
		LastName:  "Marx",
		Class:     "9C",
		Subject:   "Physics",
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	trimmedValue := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id := strings.TrimSuffix(trimmedValue, "/")
	teacherID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err", err)
		return
	}


	teachersList := make([]Teacher, 0, len(teachers))

	for _, teacher := range teachers {
		// teachersList = append(teachersList, teacher)
		if teacher.ID ==  teacherID{
			teachersList = append(teachersList, teacher)
		}
	}

	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachersList),
		Data:   teachersList,
	}
	//set content type
	w.Header().Set("Content-Type", "application/json")
	//encode data to json
	json.NewEncoder(w).Encode(response)
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Get method on /students"))
		fmt.Println("Get method on /students")
	case http.MethodPost:
		//Parse form data x-www-form-urlencoded
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte("error parsing form data"))
			return
		}
		fmt.Println("Form data:", r.Form)
		response := make(map[string]interface{})
		for key, value := range r.Form {
			response[key] = value[0]
		}
		fmt.Println("processed response:", response)
		w.Write([]byte("Post method on /students"))

		//Raw Body read and stored in a struct
		body, err := io.ReadAll(r.Body) //26:57
		if err != nil {
			log.Fatalln("error reading body", err)
			return
		}
		defer r.Body.Close()
		fmt.Println("Read Body:", string(body))
		var user1 user
		fmt.Println("user1:", user1)

		//see all what we get in request i.e. r
		fmt.Println("body:", r.Body)
		fmt.Println("form:", r.Form)
		fmt.Println("header:", r.Header)
		fmt.Println("context:", r.Context())
		fmt.Println("content length:", r.ContentLength)
		fmt.Println("host:", r.Host)
		fmt.Println("method:", r.Method)
		fmt.Println("protocol:", r.Proto)
		fmt.Println("remote address:", r.RemoteAddr)
		fmt.Println("request URI:", r.RequestURI)
		fmt.Println("tls:", r.TLS)
		fmt.Println("trailer:", r.Trailer)
		fmt.Println("transfer encoding:", r.TransferEncoding)
		fmt.Println("request url:", r.URL)
		fmt.Println("user agent:", r.UserAgent())
		fmt.Println("port:", r.URL.Port())
	case http.MethodPut:
		w.Write([]byte("Put method on /students"))
		fmt.Println("Put method on /students")
	case http.MethodDelete:
		w.Write([]byte("Delete method on /students"))
		fmt.Println("Delete method on /students")
	}
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Write([]byte("inside /exec - post method"))
		query := r.URL.Query()
		fmt.Println("Query Received:", query)
		fmt.Println("name", query.Get("name"))
		fmt.Println("name", query.Get("age"))
	}
}
