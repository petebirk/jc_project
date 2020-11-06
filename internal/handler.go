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
	"fmt"
	"os"
	"time"
)

var (
	shutdownInitiated bool = false
)

func HandleShutdown(res http.ResponseWriter, req *http.Request) {
	if (shutdownInitiated) {
		res.WriteHeader(http.StatusServiceUnavailable)
	} else {
		log.Printf("handleShutdown")
		shutdownInitiated = true;
		for (RequestsPending()) {
			time.Sleep(1 * time.Second)
		}
		res.WriteHeader(http.StatusOK)
		go os.Exit(2);
	}
}

func HandleStats(res http.ResponseWriter, req *http.Request) {
	if (shutdownInitiated) {
		res.WriteHeader(http.StatusServiceUnavailable)
	} else {
		log.Printf("handleStats")
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintf(res, "{ \"total\" : %d, \"average\" : %s }", TotalRequests(), AverageTime().String())
	}
}

func HandleHashGet(res http.ResponseWriter, req *http.Request, requestId int) {
	log.Printf("handleHashGet: requestId=%d", requestId)
	res.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	if (shutdownInitiated) {
		res.WriteHeader(http.StatusServiceUnavailable)
	} else {
		if (!ValidatePositiveInteger(requestId)) {
			log.Printf("RequestId validation failed.");
			res.WriteHeader(http.StatusBadRequest)
			return;
		} 
		password := GetHash(requestId);
		if (password == "") {
			res.WriteHeader(http.StatusNotFound)
		} else {
			res.WriteHeader(http.StatusOK)
			fmt.Fprintf(res, "%s\n", password)
		}
	}
}

func HandleHashPost(res http.ResponseWriter, req *http.Request, password string) {
	defer Track(RunningTime("HandleHashPost"))
	if (shutdownInitiated) {
		res.WriteHeader(http.StatusServiceUnavailable)
	} else {
		if (!ValidateString(password)) {
			log.Printf("Password validation failed.");
			res.WriteHeader(http.StatusBadRequest)
			return;
		}
		if (password == "") {
			log.Printf("handleHashPost: password is null. Sending BAD REQUEST.")
			res.WriteHeader(http.StatusBadRequest)
		} else {
			key := SetHash(password)
			log.Printf("handleHashPost returns requestId=%d.", key)
			res.Header().Set("Content-Type", "text/plain; charset=UTF-8")
			res.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(res, "%d\n", key)
		}
	}
}
