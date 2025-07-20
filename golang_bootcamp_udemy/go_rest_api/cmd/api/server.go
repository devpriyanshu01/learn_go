package main

import (
	"crypto/tls"
	"os"

	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error:---------", err)
		return
	}

	PORT := os.Getenv("API_PORT")
	fmt.Println("printing port no", PORT)
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer db.Close()

	cert := "cert.pem"
	key := "key.pem"

	mux := router.MainRouter()

	// mux.HandleFunc("/", handlers.RootHandler)

	// mux.HandleFunc("/teachers/", handlers.TeacherHandler)

	// mux.HandleFunc("/students/", handlers.StudentHandler)

	// mux.HandleFunc("/execs/", handlers.ExecHandler)

	//Configuring TLS handshake begins.
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Second*30)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	jwtMiddleware := mw.MiddlewaresExcludedPaths(mw.JWTMiddleware, "/execs/login")
	secureMux := jwtMiddleware(mw.SecurityHeaders(mux))
	//create custom server
	server := &http.Server{
		Addr: PORT,
		// Handler: mux,
		Handler: secureMux,
		// Handler:   rl.Middleware(mw.SecurityHeaders(mux)),
		
		// Handler: mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux)))))),

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

	fmt.Println("server is running on port:", PORT)
	// err := http.ListenAndServe(port, nil)
	err = server.ListenAndServeTLS(cert, key)
	//configuring tls handshake ends.
	if err != nil {
		fmt.Println("error:----------", err)
		log.Fatalln("Error starting the server", err)
	}
}
