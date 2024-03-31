package game

import (
	"fmt"
	"log"
	"strings"
)

type Player struct {
	Position Coordinate
}

func (g *Game) Move(nx int, ny int) {
	if nx >= 0 && ny >= 0 && nx < LEN && ny < LEN && g.Board[ny][nx] == 0 {
		g.Player.Position = Coordinate{nx, ny}
	} else {
		fmt.Printf("Invalid move to position (%d, %d)\n", nx, ny)
	}
}

func (g *Game) checkComplete() int {
	// TODO: refactor to have checkComplete return bool and a separate function askRiddle to perform other logic
	if g.Player.Position.X == LEN-1 || g.Player.Position.Y == LEN-1 {
		fmt.Println("Completed a room! Answer this riddle to move on to a new room")
		riddle, err := g.Narrator.GetRiddle()
		if err != nil {
			log.Fatal(err)
			return -1
		}
		fmt.Println(riddle)
		correct := false
		tries := 0
		for !correct {
			var answer string
			fmt.Scanln(&answer)
			resp, err := g.Narrator.SolveRiddle(riddle, answer)
			if err != nil {
				log.Fatal(err)
				return -1
			}
			if strings.Contains(resp, "True") {
				correct = true
			} else {
				fmt.Println("Incorrect answer, try again")
				tries++
				if tries > 2 {
					fmt.Println("You gave an incorrect answer three times. Game over.")
					return -1
				}
			}
		}
		g.Player.Position = Coordinate{1, 0}
		g.initBoard()
		return 0
	}
	return 1
}
