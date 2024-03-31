package game

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/teejas/gogame/narrator"
)

const LEN = 11

type Coordinate struct {
	X int
	Y int
}

type Game struct {
	Player   *Player
	Board    [LEN][LEN]int
	Narrator *narrator.Narrator
}

func (g *Game) initBoard() {
	for i := range g.Board {
		for j := range g.Board[i] {
			g.Board[i][j] = 1
		}
	}

	stack := []Coordinate{{g.Player.Position.X, g.Player.Position.Y + 1}}
	for len(stack) > 0 {
		x, y := stack[len(stack)-1].X, stack[len(stack)-1].Y
		g.Board[y][x] = 0

		neighbors := []Coordinate{
			{x + 2, y}, {x - 2, y}, {x, y + 2}, {x, y - 2},
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(neighbors), func(i, j int) {
			neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
		})

		found := false
		for _, n := range neighbors {
			nx, ny := n.X, n.Y
			if nx > 0 && nx < LEN && ny > 0 && ny < LEN && g.Board[ny][nx] == 1 {
				g.Board[ny][nx] = 0
				g.Board[y+(ny-y)/2][x+(nx-x)/2] = 0
				stack = append(stack, Coordinate{nx, ny})
				found = true
				break
			}
		}

		if !found {
			stack = stack[:len(stack)-1]
		}
	}
	g.Board[0][1] = 0
	g.Board[LEN-1][LEN-2] = 0
}

func (g Game) printBoard() {
	g.Board[g.Player.Position.Y][g.Player.Position.X] = 2
	for row := range g.Board {
		fmt.Println(g.Board[LEN-row-1])
	}
}

func (g *Game) play() {
	defer func() {
		_ = keyboard.Close()
	}()

	// Capture arrow key input
	fmt.Println("Press arrow keys (ESC to exit):")
	for {
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}
		x, y := g.Player.Position.X, g.Player.Position.Y

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		if key == keyboard.KeyArrowUp {
			g.Move(x, y+1)
		} else if key == keyboard.KeyArrowDown {
			g.Move(x, y-1)
		} else if key == keyboard.KeyArrowLeft {
			g.Move(x-1, y)
		} else if key == keyboard.KeyArrowRight {
			g.Move(x+1, y)
		} else if key == keyboard.KeyEsc {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Printf("Char: %q, Key: %v\r\n", char, key)
		}
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err = cmd.Run(); err != nil {
			log.Fatal(err)
		}
		if err := keyboard.Close(); err != nil {
			log.Fatal(err)
		}
		if g.checkComplete() == -1 {
			break
		}
		g.printBoard()
	}
}

func StartGame() {
	newPlayer := new(Player)
	newPlayer.Position = Coordinate{1, 0}
	newNarrator, err := narrator.NewNarrator()
	if err != nil {
		log.Fatal(err)
	}
	newGame := Game{
		Player:   newPlayer,
		Narrator: newNarrator,
	}
	newGame.initBoard()
	newGame.printBoard()
	newGame.play()
}
