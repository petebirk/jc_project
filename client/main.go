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

func main() {
	log.Printf("Client Started")	
	DevelopersApiRequestHashOpts opts = new DevelopersApiRequestHashOpts()
	opts.Password = "angryMonkey";		
	DevelopersApiService.RequestHash(nil, opts);

}
