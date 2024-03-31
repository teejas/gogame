package main

import (
	"fmt"
	"math/rand"
)

const LEN = 10

type Coordinate struct {
	X int
	Y int
}

type Game struct {
	PlayerPosition Coordinate
	Board          [LEN][LEN]int
}

func (g *Game) initBoard() {
	for i := range g.Board {
		for j := range g.Board[i] {
			g.Board[i][j] = 1
		}
	}

	stack := []Coordinate{g.PlayerPosition}
	for len(stack) > 0 {
		x, y := stack[len(stack)-1].X, stack[len(stack)-1].Y
		g.Board[x][y] = 0

		neighbors := []Coordinate{
			{x + 2, y}, {x - 2, y}, {x, y + 2}, {x, y - 2},
		}

		rand.Shuffle(len(neighbors), func(i, j int) {
			neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
		})

		found := false
		for _, n := range neighbors {
			nx, ny := n.X, n.Y
			if nx > 0 && nx < LEN && ny > 0 && ny < LEN && g.Board[nx][ny] == 1 {
				g.Board[nx][ny] = 0
				g.Board[x+(nx-x)/2][y+(ny-y)/2] = 0
				stack = append(stack, Coordinate{nx, ny})
				found = true
				break
			}
		}

		if !found {
			stack = stack[:len(stack)-1]
		}
	}
}

func (g Game) printBoard() {
	g.Board[g.PlayerPosition.X][g.PlayerPosition.Y] = 2
	for row := range g.Board {
		fmt.Println(g.Board[LEN-row-1])
	}
}

func main() {
	game := Game{
		PlayerPosition: Coordinate{1, 0},
	}
	game.initBoard()
	game.printBoard()
}
