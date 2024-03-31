package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/teejas/gogame/game"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	game.StartGame()
}
