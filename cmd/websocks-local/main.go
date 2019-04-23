package main

import (
	"flag"
	"github.com/yangqinjiang/websocks/core"
	"net"

	"net/url"

	"github.com/juju/loggo"
)

var localAddr string
var serverURL string
var logLevel = loggo.INFO
var debug bool

var logger = loggo.GetLogger("local")

func main() {
	flag.StringVar(&serverURL, "u", "ws://localhost:23333/websocks", "server url")
	flag.StringVar(&localAddr, "l", ":10801", "local listening port")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	if debug {
		logLevel = loggo.DEBUG
	}

	logger.SetLogLevel(logLevel)
	logger.Infof("Log level %s", logger.LogLevel().String())

	u, err := url.Parse(serverURL)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}

	lAddr, err := net.ResolveTCPAddr("tcp", localAddr)
	if err != nil {
		logger.Errorf(err.Error())
	}

	local := core.Local{
		LogLevel:   logLevel,
		ListenAddr: lAddr,
		URL:        u,
	}

	err = local.Listen()
	if err != nil {
		logger.Errorf(err.Error())
	}

}
