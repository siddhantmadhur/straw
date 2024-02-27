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
                addProjectFromArgs(args[2:])
            case "--version", "-v":
                fmt.Printf("Straw \nby Siddhant Madhur \nVersion: %s\n", version)
            case "--remove", "-r":
                err := removeProject(args[2])
                if err != nil {
                    fmt.Println("There was an error in removing directory.")
                    os.Exit(1)
                }else{
                    fmt.Println("Successfuly removed directory")
                    os.Exit(0)
                }
            default:
                fmt.Printf("Command %s not recognized.\n", args[1])
        } 
        return;
    }
}


