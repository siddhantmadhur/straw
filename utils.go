package main

import (
	"encoding/json"
	"os"
)


type Entry struct {
    Name    string `json:"name"`
    Dir     string `json:"dir"`
} 


func getEntries () ([]Entry, error) {

    data, err := os.ReadFile("data.json")    
    if err != nil {
        return []Entry{}, err
        //TODO: if error make file
    }
    var entries []Entry
    err = json.Unmarshal(data, &entries)

    if err != nil {
        return []Entry{}, err
        //TODO: if error make file
    }
    return entries, nil
}
