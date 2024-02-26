package main

import (
	"fmt"
	"os"
)

var version = "dev" 

func main() {
    args := os.Args
    entries := getEntries()
    if len(args) == 1 {
        runBubbletea(entries)    
        return;
    } else {
        switch(args[1]) {
            case "--add", "-a":
                addProject(args[2:])
            case "--version", "-v":
                fmt.Printf("Straw \nby Siddhant Madhur \nVersion: %s\n", version)

        } 
        return;
    }
}


