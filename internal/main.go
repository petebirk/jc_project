/*
 * JumpCloud Project API
 *
 * This is a simple API for JumpCloud Project
 *
 * API version: 1.0.0
 * Contact: birk.pete@gmail.com
 */
package main

import (
	"net/http"
	"log"
)

type App struct {
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
    head, req.URL.Path = ShiftPath(req.URL.Path)

	if head == "api/v1" {
        Route(res, req)
        return
    }

	http.Error(res, "Not Found", http.StatusNotFound)
}

func main() {
	log.Printf("Server started on port 8080")			

    a := &App{}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: a,
	}

    srv.ListenAndServe()
}