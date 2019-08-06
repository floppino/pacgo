package elements

import (
	"flag"
)

var (
	// ConfigFile variable
	ConfigFile = flag.String("config-file", "config.json", "path to custom configuration file")
	// MazeFile variable
	MazeFile = flag.String("Maze-file", "Maze01.txt", "path to a custom Maze file")
)

// Players is the Player character \o/
type Players struct {
	Row int
	Col int
}

// Player variable
var Player Players

// Ghost is the enemy that chases the Player :O
type Ghost struct {
	Row int
	Col int
}

// Ghosts variable
var Ghosts []*Ghost

// Config holds the emoji configuration
type Config struct {
	Player   string `json:"Player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
	UseEmoji bool   `json:"use_emoji"`
}

// Cfg variable
var Cfg Config

// Maze variable
var Maze []string

// Score variable
var Score int

// NumDots variable
var NumDots int

// Lives variable
var Lives = 1
