package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/achequisde/guessagram/def"
	"github.com/achequisde/guessagram/helpers"
)

//go:embed default_words.txt
var contents string

var quit chan int
var guess chan string

func main() {
	f, err := os.ReadFile("words.txt")
	if err != nil {
		fmt.Println("Using DEFAULT words...")
	} else {
		fmt.Println("Using LOCAL words...")
		contents = string(f)
	}

	wordList := strings.Split(strings.TrimSpace(contents), "\n")
	randomWord := helpers.PickRandomItem(wordList)
	wordSlice := strings.Split(randomWord, "")
	anagram := strings.Join(helpers.Shuffle(wordSlice), "")

	g := &def.GameState{
		End:     false,
		Tries:   1,
		Word:    randomWord,
		Anagram: anagram}

	startGame(g)
}

func startGame(g *def.GameState) {
	fmt.Println("Game start!")

	quit = make(chan int)
	guess = make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		quit <- 0
	}()

	go getUserInput(g)

	for {
		select {
		case <-quit:
			fmt.Println("Timed out!")
			os.Exit(0)
		case m := <-guess:
			g.Compare(m)

			if !g.End {
				g.Tries++
				fmt.Println("Wrong. Try again.")

				go getUserInput(g)
			} else {
				fmt.Println(g.MakeWinningMessage())
				os.Exit(0)
			}
		}
	}
}

func getUserInput(g *def.GameState) {
	fmt.Printf("Guess: %s\n", def.Cyan(g.Anagram))
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	guess <- strings.TrimSpace(text)
}
