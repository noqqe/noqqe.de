package main

import (
    "log"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

// Tiny little configuration struct
type Config struct {
  hugocmd string
  homedir string
  sammelsuriumdir string
  rvocmd string
}

// parse config from hugo repo
func parseConfig(configpath string) Config {

  filename, _ := filepath.Abs(configpath)
  yamlFile, err := ioutil.ReadFile(filename)

  if err != nil {
      panic(err)
  }

  var config Config

  err = yaml.Unmarshal(yamlFile, &config)
  if err != nil {
      panic(err)
  }

  log.Println(config)

  return config
}

