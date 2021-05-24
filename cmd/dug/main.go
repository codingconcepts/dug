package main

import (
	"flag"
	"log"
	"os"

	"github.com/codingconcepts/dug/pkg/parser"
)

func main() {
	log.SetFlags(0)

	config := flag.String("config", "", "the absolute or relative path of the sources configuration file")
	flag.Parse()

	if *config == "" {
		flag.Usage()
		os.Exit(2)
	}

	connectors, err := parser.ParseConnectors(*config)
	if err != nil {
		log.Fatalf("error parsing connectors: %v", err)
	}

	for _, c := range connectors {
		if err = c.Connect(); err != nil {
			log.Fatalf("error connecting: %v", err)
		}

		columns, err := c.Fetch()
		if err != nil {
			log.Fatalf("error fetching: %v", err)
		}

		log.Printf("%+v", c)
		for _, column := range columns {
			log.Printf("\t%+v", column)
		}
	}
}
