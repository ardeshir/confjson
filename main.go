package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Enable bool
	Path   string
	Name   string
}

func main() {

   file, _ := os.Open("conf.json")
   defer file.Close()

   decode := json.NewDecoder(file)
   conf := config{}
   err := decoder.Decode(&conf)
   if err != nil {
	fmt.Println("Error:", err)
   }

   fmt.Println(conf.Path)
   fmt.Println(conf.Name)
   fmt.Println(conf.Enabled)
}

	
