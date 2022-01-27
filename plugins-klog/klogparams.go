package main

import (
	goflag "flag"

	"github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	"k8s.io/klog"
)

var (
	str = "hello world"
	//	v       = 3
	rootCmd = &cobra.Command{
		Use:   "echo",
		Short: "use klog with cobra",
		Long:  "Use klog together with cobra.",
	}
)

func init() {
	//	rootCmd.Flags().SortFlags = false
	rootCmd.AddCommand(
		RunCmd(),
	)

	klog.InitFlags(nil)
	goflag.Set("v", "3")
	goflag.Parse()
	//pflag.MarkHidden("v")
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func RunCmd() *cobra.Command {
	runcmd := &cobra.Command{
		Use:   "run",
		Short: "run command",
		Long:  "Run command.",
		Run: func(cmd *cobra.Command, args []string) {
			klog.V(1).Infof("echo1=%v", str)
			klog.V(2).Infof("echo2=%v", str)
			klog.V(3).Infof("echo3=%v", str)
			klog.V(4).Infof("echo4=%v", str)
			greeter_main(str)
		},
	}

	//	runcmd.Flags().SortFlags = false
	runcmd.Flags().StringVar(&str, "str", str, "string to print")
	//	runcmd.Flags().IntVar(&v, "v", v, "number for the log level verbosity")

	return runcmd
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		klog.Fatalf("root cmd execute failed, err=%v", err)
	}
}
