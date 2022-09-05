package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var endpoint = flag.String(
	"endpoint",
	"myserver.aws.com",
	"The server this app will contact",
)

/*
Custom flag type
*/
type Value interface {
	String() string
	Set(string) error
}

type URLValue struct {
	URL *url.URL
}

func (u URLValue) String() string {
	if u.URL != nil {
		return u.URL.String()
	}

	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil

}

var u = &url.URL{}

// Basic flag error handling
// var (
// 	useProd = flag.Bool("prod", false, "Use a production endpoint")
// 	useDev  = flag.Bool("dev", false, "Use a development endpoint")
// 	help    = new(bool)
// )

// func init() {
// 	// flag.Var(&URLValue{u}, "url", "URL to parse")
// 	flag.BoolVar(help, "help", false, "Display help text")
// 	flag.BoolVar(help, "h", false, "Display help text (shorthand)")
// }

type Config struct {
	Name string
	Help bool
}

func main() {
	var cfg Config

	flag.StringVar(&cfg.Name, "name", "", "name of the person")
	flag.BoolVar(&cfg.Help, "help", false, "display the help text")
	flag.BoolVar(&cfg.Help, "h", false, "display the help text")
	flag.Parse()

	// fmt.Println("server endpoint is: ", *endpoint)

	// if reflect.ValueOf(*u).IsZero() {
	// 	panic("did not pass a URL")
	// }

	// fmt.Printf(`{scheme: %q, host:%q, path:%q}`, u.Scheme, u.Host, u.Path)

	if cfg.Help {
		flag.PrintDefaults()
		return
	}

	switch {
	case len(cfg.Name) == 0:
		log.Println("Error: --name of the person is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// switch {
	// case *useProd && *useDev:
	// 	log.Println("Error: --prod and --dev cannot both be set")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// case !(*useProd || *useDev):
	// 	log.Println("Error: either --prod or --dev must be set")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	fmt.Println(cfg.Name)

}
