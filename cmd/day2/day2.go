package main

import (
	"aoc2023/pkg/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Set struct {
	green int
	red   int
	blue  int
}

func (s Set) String() string {
	return fmt.Sprintf("%d green, %d red, %d blue", s.green, s.red, s.blue)
}

type Game struct {
	id   int
	sets []Set
}

func (g Game) String() string {
	return fmt.Sprintf("Game ID: %d, Sets: %v", g.id, g.sets)
}

func parseLine(line string) Game {
	// Initiate the game object
	var game Game
	var err error

	// Get the Game ID
	gameIdRegex := regexp.MustCompile(`^Game ([0-9]{1,3}): `)
	match := gameIdRegex.FindStringSubmatch(line)
	if len(match) != 2 {
		panic("The line does not contain a game ID")
	}
	game.id, err = strconv.Atoi(match[1])
	utils.CheckError(err)

	// Get the sets
	line = gameIdRegex.ReplaceAllString(line, "")
	setStrings := strings.Split(line, "; ") // Array: ["2 red, 6 green, 12 blue", "3 red, 7 green, 13 blue"]
	for _, setString := range setStrings {  // String: "2 red, 6 green, 12 blue"
		catch := strings.Split(setString, ", ") // Array: ["2 red", "6 green", "12 blue"]
		var set Set
		for _, catchItem := range catch {
			tokens := strings.Split(catchItem, " ") // Array: ["2", "red"]
			count, _ := strconv.Atoi(tokens[0])     // Int: 2
			color := tokens[1]                      // String: "red"
			switch color {
			case "red":
				set.red = count
			case "green":
				set.green = count
			case "blue":
				set.blue = count
			}
		}
		game.sets = append(game.sets, set)
	}
	return game
}

func main() {
	// Read the input file
	dir, _ := os.Getwd()
	gamesFile, _ := os.ReadFile(dir + "/cmd/day2/input.txt")
	lines := strings.Split(string(gamesFile), "\n")
	games := make([]Game, len(lines))
	// Parse all games to a slice of Game objects
	for _, line := range lines {
		games = append(games, parseLine(line))
	}

	// This is the bag that we should test against
	var bag Set = Set{
		red:   12,
		green: 13,
		blue:  14,
	}

	fails := make([]int, 0)

	// Loop through all games
	for _, game := range games {
		for _, set := range game.sets {
			if set.blue > bag.blue || set.red > bag.red || set.green > bag.green {
				fails = append(fails, game.id)
				break
			}
		}
	}

	validGames := make([]Game, 0)
	for _, game := range games {
		// Check if the game is in the fails list
		found := false
		for _, fail := range fails {
			if game.id == fail {
				found = true
				break
			}
		}
		if found {
			continue
		}
		// Game is not included in list, so it's valid
		validGames = append(validGames, game)
	}

	sum := 0
	for _, game := range validGames {
		sum += game.id
	}
}
