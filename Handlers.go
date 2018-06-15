package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// RecieveMessageHandler recieves incoming SMS messages from Twilio
func RecieveMessageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		httpErrorHandler(w, err)
		return
	}

	accountSid := r.FormValue("AccountSid")
	if accountSid != Constants.AccountSid {
		httpErrorHandlerWithCode(w, fmt.Errorf("provided AccountSid (%s) does not match this application's", accountSid), http.StatusUnauthorized)
		return
	}

	body := r.FormValue("Body")
	if len(body) == 0 {
		httpErrorHandler(w, fmt.Errorf("provided Body is empty"))
		return
	}

	from := r.FormValue("From")
	if len(from) == 0 {
		httpErrorHandler(w, fmt.Errorf("provided From is empty"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hey there! I just texted back."))
}

func httpErrorHandlerWithCode(w http.ResponseWriter, err error, code int) {
	log.Println(err)

	w.WriteHeader(code)
	w.Write([]byte("Whoops! Our system is in a 🥒!"))
}

func httpErrorHandler(w http.ResponseWriter, err error) {
	log.Println(err)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Whoops! Our system is in a 🥒!"))
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string

	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	// Return the request as a string
	return strings.Join(request, "\n")
}
