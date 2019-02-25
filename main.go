package main

import (
	"errors"
	"flag"
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	_ "xzj.com/wt/routers"
)

func main() {
	var httpport = ""
	kubeconfigPath := flag.String("kubeconfig", filepath.Join(homeDir(), ".kube", "config"), "kubeconfig绝对路径")
	context := flag.String("context", "", "集群context名称")
	port := flag.String("port", "", "监听端口号")
	flag.Parse()

	beego.Informational("加载kubeconfig的路径", *kubeconfigPath)
	beego.Informational("集群context名称", *context)

	if *context == "" {
		panic(errors.New("缺少集群context名称"))
	}
	if *port != "" {
		httpport = *port
	} else {
		httpport = beego.AppConfig.String("httpport")
	}
	beego.Informational("端口号", *port)

	beego.AppConfig.Set("kubeconfig", *kubeconfigPath)
	beego.AppConfig.Set("context", *context)

	beego.Run(":" + httpport)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

