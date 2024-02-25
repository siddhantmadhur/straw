package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
)


func addProject(args []string) {
    directory, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    globalEntries = append(globalEntries, Entry{
        Name: strings.Join(args, " "),
        Dir: strings.Replace(directory, homeDir, "~", 1),
    })
    reqBodyBytes := new(bytes.Buffer)
    json.NewEncoder(reqBodyBytes).Encode(globalEntries)
    os.WriteFile(homeDir + "/.bubble-finder", reqBodyBytes.Bytes() , 0744)
}
