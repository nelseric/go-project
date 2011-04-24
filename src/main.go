package main

import (
    board   "board"
    os      "os"
    fmt     "fmt"
    strconv "strconv"
)

func main() {
    var size = 7
    var connects = 4
    if len(os.Args) >= 3 {
        size_r := os.Args[1]
        connects_r := os.Args[2]
        size_i, err_s := strconv.Atoi(size_r)
        con_i, err_c := strconv.Atoi(connects_r)
        if err_s != nil {
            // handle error
            fmt.Println(err_s)
        } else if err_c != nil {
            // handle error
            fmt.Println(err_c)
        } else {
            if size_i >= 5 && size_i <= 20 && con_i >= 4 && con_i <= 8 {
                size = size_i
                connects = con_i
            } else {
                fmt.Println("Size or number of connects out of bounds (size 5-20, connects 4-8)")
            }
        }
    } else {
        fmt.Println("Usage: connectfour <size of board> <number of connects>\nUsing default size 7x7 with 4 connects.")
    }

    fmt.Printf("%d size board, %d connects\n", size, connects)
    b := board.NewBoard(size, connects)

    playing := true
    turn := 1
    for playing {
        b.Print()
        fmt.Printf("%s's turn. Please select a column to play.\n", (*b.Players())[turn])
        var col int
        _, err := fmt.Scanf("%d", &col)
        if err != nil || col <= 0 || col > b.Size() {
            fmt.Println("Please enter a valid number.")
        } else {
            fmt.Printf("Playing %d\n", col)
            b.Place(col-1, turn)
            if turn == 1 {
                turn = 2
            } else {
                turn = 1
            }
        }
        won, winner := b.Check()
        if won {
            b.Print()
            fmt.Printf("Congrations %s. You Won!\n", (*b.Players())[winner])
            playing = false
        }
    }
}
