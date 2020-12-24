package main

import (
	"flag"

	"github.com/wmentor/kv"
	"github.com/wmentor/log"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/langmod/controller"
)

func init() {
}

func main() {

	var dataDir string
	var listenAddr string
	flag.StringVar(&dataDir, "data", "./data", "-data dir")
	flag.StringVar(&listenAddr, "listen", "127.0.0.1:9000", "listen address like 127.0.0.1:9000")
	flag.Parse()

	if _, err := kv.Open("path=" + dataDir + " global=1"); err != nil {
		panic(err)
	}
	defer kv.Close()

	serv.File("/favicon.png", "./htdocs/favicon.png")

	serv.LoadTemplates("./templates")

	if err := serv.Start(listenAddr); err != nil {
		log.Errorf("start server failed: %s", err.Error())
	}
}
