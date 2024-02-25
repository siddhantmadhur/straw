package main

import (
	"os"
)


func main() {
    args := os.Args
    entries := getEntries()
    if len(args) == 1 {
        runBubbletea(entries)    
        return;
    } else {
        switch(args[1]) {
            case "add": 
                addProject(args[2:])
                break;  
        } 
        return;
    }
}


