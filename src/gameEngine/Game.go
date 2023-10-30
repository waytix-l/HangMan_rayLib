package gameEngine

type Game struct {
	Stage         int
	Random_number int
	Mot_secret    []string
	underscore_word string
	Content       []byte
	Err           error
	Lines         []string
	Tableau_rune  []rune
	perdu         bool
	mot           string
	win           bool
	attemptsLeft  int
	guesses       []string
	guess         string
}