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
	"time"
	"fmt"
)

var (
	totalTime time.Duration
	numRequests int32 = 0
)

func TotalRequests() int32 {
	return numRequests
}

func AverageTime() time.Duration {
	totalTimeMicroseconds := totalTime.Microseconds();
	averageMicroseconds := int64(totalTimeMicroseconds) / int64(numRequests);
	return time.Duration(averageMicroseconds * 1000)
}

func RunningTime(s string) (string, time.Time) {
    return s, time.Now()
}

func Track(s string, startTime time.Time) {
	endTime := time.Now()
	numRequests++
	currentTime := endTime.Sub(startTime);
	totalTime += currentTime;
	fmt.Print("Total Requests: ", numRequests, " CurrentTime: ", currentTime, " TotalTime: ", totalTime, " AverageTime: ", AverageTime(), "\n")
}