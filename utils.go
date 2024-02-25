package main

import (
	"encoding/json"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
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

type editorFinishedMsg struct{ err error }

func openTmux(name string, dir string) tea.Cmd {
    c := exec.Command("tmux", "new", "-s", name, "-c", dir)
    
    return tea.ExecProcess(c ,func(err error) tea.Msg {
        return tea.Quit()
    })
}
