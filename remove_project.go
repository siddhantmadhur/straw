package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func removeProject(directory string) error {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return err
    }
    directory = strings.Replace(directory, homeDir, "~", 1)
    var tempEntries []Entry
    localEntries := getEntries()
    for _, entry := range localEntries {
        if entry.Dir != directory {
            tempEntries = append(tempEntries, entry)
        } 
    }
    globalEntries = tempEntries
    reqBodyBytes := new(bytes.Buffer)
    json.NewEncoder(reqBodyBytes).Encode(tempEntries)
    err = os.WriteFile(homeDir + "/.straw", reqBodyBytes.Bytes() , 0744)
    
    return err 
}
