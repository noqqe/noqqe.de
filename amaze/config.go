package main

import (
    "log"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

// Tiny little configuration struct
type Config struct {
  Amaze struct {
    Hugocmd         string `yaml:"hugocmd"`
    Rvocmd          string `yaml:"rvocmd"`
    Homedir         string `yaml:"homedir"`
    Sammelsuriumdir string `yaml:"sammelsuriumdir"`
    Deploycmd       string `yaml:"deploycmd"`
  }
}

// parse config from hugo repo
func parseConfig(configpath string) Config {

  filename, _ := filepath.Abs(configpath)
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
      log.Panic(err)
  }

  // basically read everythin in that file
  // that results in a map[] containing interfaces{}
  // conf := make(map[interface{}]interface{})
  conf := Config{}

  err = yaml.Unmarshal(yamlFile, &conf)
  if err != nil {
      log.Panic(err)
  }

  return conf
}

