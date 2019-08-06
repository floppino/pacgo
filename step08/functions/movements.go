package functions

import (
	"fmt"
	"math/rand"
	"os"

	v "github.com/pacgo/step08/elements"
)

// MoveCursor print the cursor
func MoveCursor(Row, Col int) {
	if v.Cfg.UseEmoji {
		fmt.Printf("\x1b[%d;%df", Row+1, Col*2+1)
	} else {
		fmt.Printf("\x1b[%d;%df", Row+1, Col+1)
	}
}

// ReadInput reads the keyboard input
func ReadInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

// MakeMove convert the move on the map
func MakeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(v.Maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(v.Maze)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(v.Maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(v.Maze[0]) - 1
		}
	}

	if v.Maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

// MovePlayer moves Player and increment score
func MovePlayer(dir string) {
	v.Player.Row, v.Player.Col = MakeMove(v.Player.Row, v.Player.Col, dir)
	switch v.Maze[v.Player.Row][v.Player.Col] {
	case '.':
		v.NumDots--
		v.Score++
		// Remove dot from the Maze
		v.Maze[v.Player.Row] = v.Maze[v.Player.Row][0:v.Player.Col] + " " + v.Maze[v.Player.Row][v.Player.Col+1:]
	}
}

// DrawDirection randomly decides ghosts' direction
func DrawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

// MoveGhosts moves the ghosts
func MoveGhosts() {
	for _, g := range v.Ghosts {
		dir := DrawDirection()
		g.Row, g.Col = MakeMove(g.Row, g.Col, dir)
	}
}
