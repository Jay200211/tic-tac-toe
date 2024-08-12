package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Empty = iota
	PlayerX
	PlayerO
)

type Board [3][3]int

type Game struct {
	board  Board
	player int
}

func NewGame() *Game {
	return &Game{
		board:  Board{},
		player: PlayerX,
	}
}
func (g *Game) DisplayBoard() {
	for _, row := range g.board {
		for _, cell := range row {
			switch cell {
			case PlayerX:
				fmt.Print("X ")
			case PlayerO:
				fmt.Print("O ")
			default:
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func (g *Game) MakeMove(row, col int) bool {
	if g.board[row][col] == Empty {
		g.board[row][col] = g.player
		return true
	}
	return false
}
func (g *Game) CheckWin() int {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] && g.board[i][0] != Empty {
			return g.board[i][0]
		}
		if g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] && g.board[0][i] != Empty {
			return g.board[0][i]
		}
	}
	// Check diagonals
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[0][0] != Empty {
		return g.board[0][0]
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[0][2] != Empty {
		return g.board[0][2]
	}
	return Empty
}
func (g *Game) AITurn() {
	// Check for winning move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == Empty {
				g.board[i][j] = PlayerO
				if g.CheckWin() == PlayerO {
					return
				}
				g.board[i][j] = Empty
			}
		}
	}
	// Check for blocking move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == Empty {
				g.board[i][j] = PlayerX
				if g.CheckWin() == PlayerX {
					g.board[i][j] = PlayerO
					return
				}
				g.board[i][j] = Empty
			}
		}
	}
	// Random move
	rand.Seed(time.Now().UnixNano())
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if g.board[row][col] == Empty {
			g.board[row][col] = PlayerO
			return
		}
	}
}
func main() {
	game := NewGame()
	for {
		game.DisplayBoard()
		var row, col int
		fmt.Printf("Player X, enter your move (row and column): ")
		fmt.Scan(&row, &col)

		if !game.MakeMove(row, col) {
			fmt.Println("Invalid move, try again.")
			continue
		}

		if winner := game.CheckWin(); winner != Empty {
			game.DisplayBoard()
			fmt.Printf("Player %v wins!\n", winner)
			break
		}

		game.AITurn()

		if winner := game.CheckWin(); winner != Empty {
			game.DisplayBoard()
			fmt.Printf("Player %v wins!\n", winner)
			break
		}
	}
}
