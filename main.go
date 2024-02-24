package main

import (
	"os/exec"
)


func main() {
    _, err := getEntries()

    if err != nil {
        panic(err)
    }

    //runBubbletea(entries)    
    
    cmd := exec.Command("cd ~/github/project_1")
    cmd.Run()
}
