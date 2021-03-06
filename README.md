go-tigertonic-yaml-config
=========================

YAML parsing extension for github.com/rcrowley/go-tigertonic.

To use this, import it for the side-effect of registering the YAML parser with
go-tigertonic.


```Go
package main

import (
	"github.com/bigkevmcd/go-tigertonic"
	// Only need this to register itself.
	_ "github.com/bigkevmcd/go-tigertonic-yaml-config"
)

var config = flag.String("config", "", "pathname of JSON configuration file")

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: example [-config=<config>]")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	// Example of parsing a configuration file.
	c := &Config{}
	if err := tigertonic.Configure(*config, c); nil != err {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", c)
}
```
