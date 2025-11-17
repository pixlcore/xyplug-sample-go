package main

import (
    "encoding/json"
    "fmt"
    "io"
    "os"

    "github.com/fatih/color"
)

func main() {
    // Read all data from STDIN
    data, err := io.ReadAll(os.Stdin)
    if err == nil && len(data) > 0 {
        var obj any
        if err := json.Unmarshal(data, &obj); err != nil {
            fmt.Fprintf(os.Stderr, "Error: failed to parse JSON from STDIN: %v\n", err)
        } else {
            pretty, _ := json.MarshalIndent(obj, "", "  ")
            fmt.Printf("Read JSON from STDIN:\n%s\n\n", string(pretty))
        }
    }

    // Print a simple green message using the dependency
    color.Green("Hello from dependency github.com/fatih/color!")

    // Always print final JSON message
    fmt.Println(`{ "xy":1, "code":0 }`)
}
