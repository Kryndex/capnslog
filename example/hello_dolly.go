package main

import (
	"github.com/coreos/pkg/capnslog"
	oldlog "log"
	"os"
)

var log = capnslog.NewPackageLogger("github.com/coreos/pkg/capnslog/cmd", "main")
var dlog = capnslog.NewPackageLogger("github.com/coreos/pkg/capnslog/cmd", "dolly")

func main() {
	rl := capnslog.MustRepoLogger("github.com/coreos/pkg/capnslog/cmd")
	rl.SetGlobalLogLevel(capnslog.INFO)
	capnslog.SetFormatter(capnslog.NewGlogFormatter(os.Stderr))

	// We can parse the log level configs from the command line
	if len(os.Args) > 1 {
		cfg, err := rl.ParseLogLevelConfig(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}
		rl.SetLogLevel(cfg)
		log.Infoln("Setting output to", os.Args[1])
	}

	// Send some messages at different levels to the different packages
	dlog.Infoln("Hello Dolly")
	dlog.Warningln("Well hello, Dolly")
	log.Errorln("It's so nice to have you back where you belong")
	dlog.Debugln("You're looking swell, Dolly")
	dlog.Verboseln("I can tell, Dolly")

	// We also have control over the built-in "log" package.
	ol := capnslog.MustRepoLogger("log")
	ol.SetGlobalLogLevel(capnslog.INFO)
	oldlog.Println("You're still glowin', you're still crowin', you're still lookin' strong")
	log.Fatalln("Dolly'll never go away again")
}
