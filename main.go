package main


func main() {
    entries, err := getEntries()

    if err != nil {
        panic(err)
    }

    runBubbletea(entries)    
      
}
