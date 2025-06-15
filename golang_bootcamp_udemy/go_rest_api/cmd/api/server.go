package main

import (
	"crypto/tls"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strings"
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

	//create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler:   rl.Middleware(mw.SecurityHeaders(mux)),
		Handler:   mw.Hpp(hppOptions)(rl.Middleware(mw.SecurityHeaders(mux))),
		// Handler:   middlewares.Cors(mux),
		// Handler: middlewares.ResponseTimeMiddleware(mux),
		// Handler: middlewares.Compression(mux),
		TLSConfig: tlsConfig,
	}

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
	// fmt.Println(r.Method)
	// if r.Method == http.MethodGet {
	// 	w.Write([]byte("get request"))
	// }
	switch r.Method {
	case http.MethodGet:
		//path parameters
		//teachers/{id}  <-- path parameter
		//teachers/?id=23&name=raman&age=20
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		pathParam := strings.TrimPrefix(path, "/")
		fmt.Println("path param:", pathParam)
		w.Write([]byte("get method /teachers handled"))

		//query parameters
		//teachers/?id=23&name=raman&age=20
		queryParameters := r.URL.Query()
		fmt.Println("printing query params received")
		fmt.Println(queryParameters.Get("id"))
		fmt.Println(queryParameters.Get("name"))
		fmt.Println(queryParameters.Get("age"))

	case http.MethodPost:
		fmt.Fprintf(w, "post method /teachers")
	}
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
