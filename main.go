package main

import (
	"fmt"
	"os"
)

var version = "0.1.0"

func main() {
    args := os.Args
    entries := getEntries()
    if len(args) == 1 {
        runBubbletea(entries)    
        return;
    } else {
        switch(args[1]) {
            case "--add":
            case "-a":
                addProject(args[2:])
                break;  
            case "--version":
            case "-v":
                fmt.Printf("Straw \nby Siddhant Madhur \nVersion: v%s\n", version)

        } 
        return;
    }
}


