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
	 "crypto/sha512"
	 "encoding/base64"
 )
 
 func HashPassword(password string) string {
	 sha512 := sha512.New()
	 bv := []byte(password);
	 sha512.Write(bv)
	 encoded := base64.URLEncoding.EncodeToString(sha512.Sum(nil))
	 return encoded
 }
 
