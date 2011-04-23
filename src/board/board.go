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
    for i := b.size - 1; i >= 0; i-- {
        for j := 0; j < b.size; j++ {
            fmt.Printf("%s ", b.players[b.pieces[j][i]])
        }
        fmt.Println()
    }
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

func (b *Board) Check() int {
    //check vertically.
    consecutive := 1
    prev := 0

    for row := 0; row < b.size; row++ {
        for col := 0; col < b.size; col++ {

            if b.pieces[col][row] == prev && prev != '_' {
                consecutive++
            } else {
                consecutive = 1
            }
            if consecutive >= b.connects {
                return b.pieces[col][row]
            }
            prev = b.pieces[col][row]

        }
    }
    /*
    	//check horizontally.
    	consecutive = 1;
    	prev = 0;

    	for (row = 0; row < size; ++row) {
    		for (col = 0; col < size; ++col) {
    			i = row * size + col;

    			if (b->pieces[i] == prev && prev != '_') {
    				consecutive++;
    			} else {
    				consecutive = 1;
    			}
    			if (consecutive >= b->connects) {
    				return b->pieces[i];
    			}
    			prev = b->pieces[i];

    		}
    	}
    	//check diagonally /.
    	for (row = size - 1; row >= b->connects; --row) {
    		for (col = 0; col < size - 1 - b->connects; col++) {
    			i = row * size + col;
    			if (b->pieces[i] != '_') {
    				int diag_mod;
    				int d;
    				char complete = 1;
    				for (diag_mod = 1; diag_mod < b->connects; ++diag_mod) {
    					d = (row - diag_mod) * size + (col + diag_mod);
    					if (b->pieces[i] != b->pieces[d]) {
    						complete = 0;
    						break;
    					}
    				}
    				if (complete) {
    					return b->pieces[i];
    				}
    			}
    		}
    	}

    	//check diagonally \.

    	for (row = size - 1; row >= b->connects; --row) {
    		for (col = size-1; col >=  b->connects-1; col--) {
    			i = row * size + col;
    			if (b->pieces[i] != '_') {
    				int diag_mod;
    				int d;
    				char complete = 1;
    				for (diag_mod = 1; diag_mod < b->connects; ++diag_mod) {
    					d = (row - diag_mod) * size + (col - diag_mod);
    					if (b->pieces[i] != b->pieces[d]) {
    						complete = 0;
    						break;
    					}
    				}
    				if (complete) {
    					return b->pieces[i];
    				}
    			}
    		}
    	}
    */
    return 0
}
