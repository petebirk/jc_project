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
	"strconv"
)

func Route(res http.ResponseWriter, req *http.Request) {
	var operation, pathVariable string
	operation, pathVariable = ShiftPath(req.URL.Path)

	switch operation {
	case "hash":
		if (pathVariable == "/") {
			switch req.Method {
			case "POST":
				req.ParseForm()
				var password string;
				for key, value := range req.Form {
					if (key == "password") {
						password = value[0];
					}
				} 

				HandleHashPost(res, req, password)
			default:
				http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		} else {
			switch req.Method {
			case "GET":
				pathVariable = pathVariable[1:];
				requestId, err := strconv.Atoi(pathVariable)
				if (err == nil) {
					HandleHashGet(res, req, requestId)
				} else {
					http.Error(res, "Bad Request", http.StatusBadRequest)
				}
			default:
				http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		}
	case "stats":
		switch req.Method {
		case "GET":
			HandleStats(res, req)
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	case "shutdown":
		switch req.Method {
		case "POST":
			HandleShutdown(res, req)
		default:
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(res, "Path Not Found", http.StatusNotFound)
	}

 }
