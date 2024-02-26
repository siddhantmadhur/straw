package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)


type Entry struct {
    Name    string `json:"name"`
    Dir     string `json:"dir"`
} 

var globalEntries []Entry

func fileExists(filename string) bool {
   info, err := os.Stat(filename)
   if os.IsNotExist(err) {
      return false
   }
   return !info.IsDir()
}

func getEntries () []Entry {

    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    data, err := os.ReadFile(homeDir + "/.straw")    
    if err != nil {
        return []Entry{}
        //TODO: if error make file
    }
    var entries []Entry
    err = json.Unmarshal(data, &entries)
    globalEntries = entries
    if err != nil {
        return []Entry{}
        //TODO: if error make file
    }
    return entries
}

type editorFinishedMsg struct{ err error }

func openTmux(name string, directory string) tea.Cmd {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }

    c := exec.Command("tmux", "new", "-A", "-s", name, "-c", homeDir + strings.Replace(directory, "~", "", 1))
    
    return tea.ExecProcess(c ,func(err error) tea.Msg {
        return tea.Quit()
    })
}
