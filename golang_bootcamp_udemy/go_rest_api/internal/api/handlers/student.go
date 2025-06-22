package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
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
		var user1 models.User
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
