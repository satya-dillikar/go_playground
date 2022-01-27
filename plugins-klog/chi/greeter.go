package main

import (
	"fmt"

	"k8s.io/klog"
)

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
	str := "chinese"
	klog.V(1).Infof("echo1=%v", str)
	klog.V(2).Infof("echo2=%v", str)
	klog.V(3).Infof("echo3=%v", str)
	klog.V(4).Infof("echo4=%v", str)
}

var Greeter greeting
