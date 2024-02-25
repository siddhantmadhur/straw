package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
    for _, entry := range globalEntries {
        if entry.Dir == strings.Replace(directory, homeDir, "~", 1) || entry.Name == strings.Join(args, " ") {
            fmt.Println("This directory or name already exists")
            os.Exit(1)
        }
    }
    globalEntries = append(globalEntries, Entry{
        Name: strings.Join(args, " "),
        Dir: strings.Replace(directory, homeDir, "~", 1),
    })
    sort.Slice(globalEntries, func (i, j int) bool {
        return globalEntries[i].Name  < globalEntries[j].Name
    })
    reqBodyBytes := new(bytes.Buffer)
    json.NewEncoder(reqBodyBytes).Encode(globalEntries)
    os.WriteFile(homeDir + "/.straw", reqBodyBytes.Bytes() , 0744)
    fmt.Println("Added successfully")
}
