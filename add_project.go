package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func addProject(name string, directory string) error { 
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    for _, entry := range globalEntries {
        if entry.Dir == strings.Replace(directory, homeDir, "~", 1) || entry.Name == name {
            fmt.Println("This directory or name already exists")
            os.Exit(1)
        }
    }
    globalEntries = append(globalEntries, Entry{
        Name: name,
        Dir: strings.Replace(directory, homeDir, "~", 1),
    })
    sort.Slice(globalEntries, func (i, j int) bool {
        return globalEntries[i].Name  < globalEntries[j].Name
    })
    reqBodyBytes := new(bytes.Buffer)
    err = json.NewEncoder(reqBodyBytes).Encode(globalEntries)
    err = os.WriteFile(homeDir + "/.straw", reqBodyBytes.Bytes() , 0744)
    return err
}


func addProjectFromArgs(args []string) {
    directory, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    err = addProject(strings.Join(args, " "), directory)
    if err != nil {
        panic(err)
    }
    fmt.Println("Added successfully")
}
