package main

import (
    "log"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

// Tiny little configuration struct
type Config struct {
  hugocmd         string  //`yaml:"hugocmd"`
  rvocmd          string  //`yaml:"rvocmd"`
  homedir         string  //`yaml:"homedir"`
  sammelsuriumdir string  //`yaml:"sammelsuriumdir"`
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
  conf := make(map[interface{}]interface{})

  err = yaml.Unmarshal(yamlFile, &conf)
  if err != nil {
      log.Panic(err)
  }

  // pick the subcomponent "section" from the map
  // which is an interface
  section := conf["amaze"]

  // convert the interface(!) "section" to a map.
  // if this doesnt happen, the indexing does not work
  s, _ := section.(map[interface{}]string)

  // its very disgusting. map[interface{}]interface{} is a bitch.
  // even if remapped
  c := Config{
  hugocmd: s["hugocmd"],
  rvocmd: s["rvocmd"],
  homedir: s["homedir"],
  sammelsuriumdir: s["sammelsuriumdir"]}

  return c
}

