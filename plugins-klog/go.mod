module satya.com/plugins-klog

go 1.17

require (
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	k8s.io/klog v1.0.0
)

require github.com/inconshreveable/mousetrap v1.0.0 // indirect

replace satya.com/plugins-klog => ./
