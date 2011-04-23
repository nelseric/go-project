package main

import (
    board   "board"
    os      "os"
    fmt     "fmt"
    strconv "strconv"
)

func main() {
    //b := board.NewBoard(8, 4)

    size := os.Args[1]

    i, err := strconv.Atoi(size)
    if err != nil {
        // handle error
        fmt.Println(s, err)
        os.Exit(2)
    }

    //b.Print()

}
