package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Props = initProps()

func initProps() props {
	p := new(props)
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Printf("Error opening config.yaml <%v>\n", err)
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		log.Fatalf("Error unmarshalling configuration file <%v>", err)
	}
	return *p
}
