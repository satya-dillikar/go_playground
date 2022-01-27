// Program in GO language to demonstrates how to use base log package.
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("init started")
}
func main() {
	// Println writes to the standard logger.
	log.Println("main started")

	/* 	// Connect to the remote SMTP server.
	   	client, err := smtp.Dial("smtp.smail.com:25")
	   	if err != nil {
	   		log.Println(err)
	   	}
	   	//client.Data() */

	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic()
	//log.Panicln("panic message")

}
