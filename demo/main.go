package main

import (
	"encoding/json"
	"log"

	"github.com/themakers/osinfo"
)

func main() {
	osi := osinfo.GetInfo()

	data, _ := json.MarshalIndent(osi, "", " ")

	log.Println(string(data))
}
