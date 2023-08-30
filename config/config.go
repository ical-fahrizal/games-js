package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/ipsusila/opt"
)

var configPathView string

// const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func init() {
	// with loggger lumberjack
	ex, err := os.Executable()
	if err != nil {
		log.Println("Error Executable: ", err)
	}
	exPath := filepath.Dir(ex) + "/"
	log.Printf("exPath : %v", exPath)
	confPath := exPath + "config.hjson"
	log.Printf("confPath : %v", confPath)

	// re-open file
	file, err := os.Open(confPath)
	if err != nil {
		confPath = "config.hjson"
	}
	defer file.Close()

	//parse configurationf file
	cfgFile := flag.String("conf", confPath, "Configuration file")
	flag.Parse()
	log.Print("masuk config")

	//load options
	config, err := opt.FromFile(*cfgFile, opt.FormatAuto)
	if err != nil {
		log.Printf("Error while loading configuration file %v -> %v\n", *cfgFile, err)
		return
	}

	//app_name config
	configPathView = config.Get("parameter").GetString("pathView", "")
	log.Printf("configPathView : %v", configPathView)

}

func GetPathView() string {
	return configPathView
}
