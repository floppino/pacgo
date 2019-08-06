package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	v "github.com/pacgo/step08/elements"
	fun "github.com/pacgo/step08/functions"
)

func loadConfig() error {
	file, err := os.Open(*v.ConfigFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v.Cfg)
	if err != nil {
		return err
	}

	return nil
}

func initialize() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
	}
}

func main() {
	flag.Parse()

	// initialize game
	initialize()
	defer cleanup()

	// load resources
	err := fun.LoadMaze()
	if err != nil {
		log.Printf("Error loading Maze: %v\n", err)
		return
	}

	err = loadConfig()
	if err != nil {
		log.Printf("Error loading configuration: %v\n", err)
		return
	}

	// process input (async)
	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := fun.ReadInput()
			if err != nil {
				log.Printf("Error reading input: %v", err)
				ch <- "ESC"
			}
			ch <- input
		}
	}(input)

	// game loop
	for {
		// process movement
		select {
		case inp := <-input:
			if inp == "ESC" {
				v.Lives = 0
			}
			fun.MovePlayer(inp)
		default:
		}

		fun.MoveGhosts()

		// process Collisions
		for _, g := range v.Ghosts {
			if v.Player.Row == g.Row && v.Player.Col == g.Col {
				v.Lives = 0
			}
		}

		// update screen
		fun.PrintScreen()

		// check game over
		if v.NumDots == 0 || v.Lives == 0 {
			if v.Lives == 0 {
				fun.MoveCursor(v.Player.Row, v.Player.Col)
				fmt.Printf(v.Cfg.Death)
				fun.MoveCursor(len(v.Maze)+2, 0)
			}
			break
		}

		// repeat
		time.Sleep(200 * time.Millisecond)
	}
}
