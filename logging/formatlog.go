// Golang program to print specific date and time
package main

import (
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc/status"
)

func main() {
	//dt := time.Now()
	/*
		// Format MM-DD-YYYY
		fmt.Println(dt.Format("01-02-2006"))

		// Format MM-DD-YYYY hh:mm:ss
		fmt.Println(dt.Format("01-02-2006 15:04:05"))

		// With short weekday (Mon)
		fmt.Println(dt.Format("01-02-2006 15:04:05 Mon"))

		// With weekday (Monday)
		fmt.Println(dt.Format("01-02-2006 15:04:05 Monday")) */

	// Include micro seconds

	logRequest := CreateRequestLogger()
	logRequest()

}

func shouldLogEndpoint(endpoint string) bool {

	// supressLoggingOfEndpoints := []string{"GetInstalledPackageSummaries",
	//	"GetAvailablePackageSummaries"}

	supressLoggingOfEndpoints := []string{"GetConfiguredPlugins"}

	//log.Printf("shouldLogEndpoint: %s", endpoint)

	suppress := true
	for i := 0; i < len(supressLoggingOfEndpoints); i++ {
		if strings.Contains(endpoint, supressLoggingOfEndpoints[i]) {
			suppress = false
			break
		}
	}
	return suppress
}

// LogRequest is a gRPC UnaryServerInterceptor that will log the API call
func CreateRequestLogger() func() (err error) {

	// Include micro seconds in timestamp
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	// log to custom file
	// LOG_FILE := "./myapp_log"
	// open log file
	// logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// Don't defer  beacause it will close immediately.
	//defer logFile.Close()

	// Set log out put and enjoy :)
	//logger.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	logger.SetFlags(log.Lshortfile | log.LstdFlags)

	return func() (err error) {

		start := time.Now()

		// Format string : [timestamp] [status code] [duration] [full path]
		// 2021-11-29 15:10:21.642313 OK 97.752µs /kubeappsapis.core.packages.v1alpha1.PackagesService/GetAvailablePackageSummaries
		FullMethod := "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetConfiguredPlugins"

		if shouldLogEndpoint(FullMethod) {
			logger.Printf("%v %s %s\n",
				status.Code(nil),
				time.Since(start),
				FullMethod)
		}

		FullMethod = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetInstalledPackageSummaries"

		if shouldLogEndpoint(FullMethod) {
			logger.Printf("%v %s %s\n",
				status.Code(nil),
				time.Since(start),
				FullMethod)
		}

		FullMethod = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetAvailablePackageSummaries"

		if shouldLogEndpoint(FullMethod) {
			logger.Printf("%v %s %s\n",
				status.Code(nil),
				time.Since(start),
				FullMethod)
		}

		FullMethod = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetConfiguredPlugins"

		if shouldLogEndpoint(FullMethod) {
			logger.Printf("%v %s %s\n",
				status.Code(nil),
				time.Since(start),
				FullMethod)
		}

		return nil
	}
}
