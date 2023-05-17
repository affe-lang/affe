package main

import (
    "fmt"
    "os"
    "os/user"
    "affe/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf(`        __  __     
       / _|/ _|      | logged in as %s
  __ _| |_| |_ ___   | 
 / _' |  _|  _/ _ \  | version 0.1
| (_| | | | ||  __/  | press ctrl+c to exit
 \__,_|_| |_| \___|  | https://affe-lang.github.io

`, user.Username);
    repl.Start(os.Stdin, os.Stdout)
}
