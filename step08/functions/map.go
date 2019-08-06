package functions

import (
	"bufio"
	"fmt"
	"os"

	v "github.com/pacgo/step08/elements"
)

// PrintScreen to show the map on the screen
func PrintScreen() {
	ClearScreen()
	for _, line := range v.Maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Printf(v.Cfg.Wall)
			case '.':
				fmt.Printf(v.Cfg.Dot)
			default:
				fmt.Printf(v.Cfg.Space)
			}
		}
		fmt.Printf("\n")
	}

	MoveCursor(v.Player.Row, v.Player.Col)
	fmt.Printf(v.Cfg.Player)

	for _, g := range v.Ghosts {
		MoveCursor(g.Row, g.Col)
		fmt.Printf(v.Cfg.Ghost)
	}

	MoveCursor(len(v.Maze)+1, 0)
	fmt.Printf("Score: %v\tLives: %v\n", v.Score, v.Lives)
}

// ClearScreen to clear the screen
func ClearScreen() {
	fmt.Printf("\x1b[2J")
	MoveCursor(0, 0)
}

// LoadMaze to loead the map
func LoadMaze() error {
	file, err := os.Open(*v.MazeFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		v.Maze = append(v.Maze, line)
	}

	for row, line := range v.Maze {
		for col, char := range line {
			switch char {
			case 'P':
				v.Player = v.Players{row, col}
			case 'G':
				v.Ghosts = append(v.Ghosts, &v.Ghost{row, col})
			case '.':
				v.NumDots++
			}
		}
	}
	return nil
}
