package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Addr string `yaml:addr`
		Port string `yaml:portt`
	}
	Tran struct {
		Interval int `yaml:interval`
	}
}

func getConfig() Config {
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var cnf Config
	err = yaml.Unmarshal(bytes, &cnf)
	if err != nil {
		panic(err)
	}
	return cnf
}
