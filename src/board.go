package board

import fmt "fmt"

type Board struct {
    pieces   [][]int
    colsize  []int
    players  map[int]string
    size     int
    connects int
}

func NewBoard(size, connects int) *Board {
    b := new(Board)
    b.connects = connects
    b.pieces = make([][]int, size)
    b.colsize = make([]int, size)
    for i := 0; i < size; i++ {
        b.pieces[i] = make([]int, size)
    }
    b.players = map[int]string{
        0:  "_",
        1:  "X",
        2:  "O",
    }
    b.size = size
    return b
}

func (b *Board) Size() int {
    return b.size
}

func (b *Board) Players() *map[int]string {
    return &b.players
}

func (b *Board) Pieces() *[][]int {
    return &b.pieces
}

func (b *Board) Print() {
    fmt.Print("┌─")
    for i := 1; i < b.size; i++ {
        fmt.Print("┬─")
    }
    fmt.Println("┐")
    for i := b.size - 1; i >= 0; i-- {
        for j := 0; j < b.size; j++ {
            fmt.Printf("│%s", b.players[b.pieces[j][i]])

        }
        fmt.Println("│")
    }
    fmt.Print("└─")
    for i := 1; i < b.size; i++ {
        fmt.Print("┴─")
    }
    fmt.Println("┘")
}

func (b *Board) Place(col, player int) bool {
    var ret bool
    if b.colsize[col] == b.size {
        ret = false
    } else {
        b.pieces[col][b.colsize[col]] = player
        b.colsize[col]++
        ret = true
    }
    return ret
}

func (b *Board) Check() (is_win bool, winner int) {
    //check vertically.
    consecutive := 1
    prev := 0

    for row := 0; row < b.size; row++ {
        for col := 0; col < b.size; col++ {

            if b.pieces[col][row] == prev && prev != 0 {
                consecutive++
            } else {
                consecutive = 1
            }
            if consecutive >= b.connects {
                return true, b.pieces[col][row]
            }
            prev = b.pieces[col][row]

        }
    }
    //check horizontally.
    consecutive = 1
    prev = 0

    for col := 0; col < b.size; col++ {
        for row := 0; row < b.size; row++ {

            if b.pieces[col][row] == prev && prev != 0 {
                consecutive++
            } else {
                consecutive = 1
            }
            if consecutive >= b.connects {
                return true, b.pieces[col][row]
            }
            prev = b.pieces[col][row]

        }
    }

    // checck diagonally /
    for col := 0; col <= b.size-b.connects; col++ {
        for row := 0; row <= b.size-b.connects; row++ {

            if b.pieces[col][row] != 0 {
                complete := true
                for dm := 1; dm < b.connects; dm++ {
                    row_d := row + dm
                    col_d := col + dm
                    if b.pieces[col][row] != b.pieces[row_d][col_d] {
                        complete = false
                        break
                    }
                }
                if complete {
                    return true, b.pieces[col][row]
                }
            }

        }
    }
    // checck diagonally \
    for col := 0; col <= b.size-b.connects; col++ {
        for row := b.size - 1; row >= b.connects-1; row-- {
            if b.pieces[col][row] != 0 {
                complete := true
                for dm := 1; dm < b.connects; dm++ {
                    row_d := row - dm
                    col_d := col + dm
                    if b.pieces[col][row] != b.pieces[col_d][row_d] {
                        complete = false
                        break
                    }
                }
                if complete {
                    return true, b.pieces[col][row]
                }
            }

        }
    }

    return false, 0
}
