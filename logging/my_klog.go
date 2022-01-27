// Program in GO language to demonstrates how to use base log package.
package main

import (
	"flag"
	"fmt"

	log "k8s.io/klog/v2"

	"k8s.io/klog/v2/klogr"
)

type myError struct {
	str string
}

func (e myError) Error() string {
	return e.str
}
func main() {
	log.InitFlags(nil)
	// By default klog writes to stderr. Setting logtostderr to false makes klog
	// write to a log file.
	//flag.Set("logtostderr", "false")
	//flag.Set("log_file", "myfile.log")

	// In this example, we want to show you that all the lines logged
	// end up in the myfile.log. You do NOT need them in your application
	// as all these flags are set up from the command line typically
	//flag.Set("logtostderr", "false")     // By default klog logs to stderr, switch that off
	flag.Set("alsologtostderr", "false") // false is default, but this is informative
	flag.Set("stderrthreshold", "FATAL") // stderrthreshold defaults to ERROR, we don't want anything in stderr
	//flag.Set("log_file", "myfile.log")   // log to a file
	flag.Parse()
	log.Info("INFO- nice to meet you")

	// Println writes to the standard logger.
	log.Warningf("WARN -main started")

	// Fatalln is Println() followed by a call to os.Exit(1)
	//log.Fatalln("fatal message")

	log.Errorln("Error - report new line")
	log.Errorf("Error - report")
	log.Flush()
	fmt.Println("------------")

	my_klogr := klogr.New().WithName("MyName").WithValues("user", "you")
	my_klogr.Info("INFO- nice to meet you")
	// Println writes to the standard logger.
	//my_klogr.Infof("INFOf -main started")

	// Fatalln is Println() followed by a call to os.Exit(1)
	//my_klogr.Fatalln("fatal message")

	my_klogr.Error("Error - report new line")
	//my_klogr.Errorf("Errorf - report")
	//my_klogr.Flush()
}
