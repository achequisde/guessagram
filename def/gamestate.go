package def

import "fmt"

type GameState struct {
	End     bool
	Tries   int
	Word    string
	Anagram string
}

func (g *GameState) Compare(guess string) {
	g.End = guess == g.Word
}

func (g *GameState) MakeWinningMessage() string {
	plural := g.Tries > 1

	var triesOrTry string

	if plural {
		triesOrTry = "tries"
	} else {
		triesOrTry = "try"
	}

	return fmt.Sprintf("You won in %d %s!", g.Tries, triesOrTry)
}
