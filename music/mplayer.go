package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maborosii/music/library"
	"github.com/maborosii/music/mp"
)

var lib *library.MusicManager
var id int = 1

// var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		if lib.Len() > 0 {
			for i := 0; i < lib.Len(); i++ {
				e, _ := lib.Get(i)
				fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
			}
		} else {
			fmt.Println("the list is empty")
		}

	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{Id: strconv.Itoa(id),
				Name: tokens[2], Artist: tokens[3], Source: tokens[4], Type: tokens[5]})
			fmt.Println("add success")
		} else {
			fmt.Println("Usage: lib add <name> <artist> <source> <type>")
		}
	case "remove":
		if len(tokens) == 3 {
			num, _ := strconv.Atoi(tokens[2])
			lib.Remove(num - 1)
			fmt.Println("delete success")
		} else {
			fmt.Println("Usage: lib remove <id>")
		}
	default:
		fmt.Println("Unrecongnized lib command: ", tokens[1])
	}

}

func hanlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Play <name>")
		return
	}

	e := lib.Find(tokens[1])

	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist")
		return
	}

	mp.PlayMusic(e.Source, e.Type)
}

func main() {
	fmt.Println(`
				11111
				22222
				33333	
	`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command ->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" && len(tokens) >= 2 {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			hanlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecongnized command: ", tokens[0])
			// continue
		}
	}
}
