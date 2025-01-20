package game

import (
	"GoTicTacToe/internal/log"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	sWidth  int // The screen width
	sHeight int // The screen height

	state    GameState    // The current game state
	gameMode int          // The game mode (1 = easy AI, 2 = hard AI, 3 = GoRythm mode)
	playing  string       // The current turn player ("O" or "X")
	player   string       // The player type ("human" or "ai")
	board    [3][3]string // The game board
	pointsO  int          // The point number for player O
	pointsX  int          // The point number for player X
	rounds   int          // The number of rounds
	win      string       // The win status

	goRythm *GoRythm // GoRythm mode game struct

	audioContext *audio.Context // The audio context for the game
	audioPlayer  *AudioPlayer   // The audio player for the game used to play the music

	countdownTime time.Time // The countdown timer
	countdown     int       // The countdown duration

	gameImage                            *ebiten.Image // The game image containing the background and symbols are drawn on it
	boardImage                           *ebiten.Image // The board grid image
	XImage, OImage                       *ebiten.Image // The symbols images
	XImageHighlighted, OImageHighlighted *ebiten.Image // The highlighted symbols images (for GoRythm mode)
	EmptyImage                           *ebiten.Image // The empty symbol image (for removing symbols in GoRythm mode)
}

type GameState int

const (
	StateMenu GameState = iota
	StateLoading
	StatePlaying
	StatePause
	StateGameOver
)

const (
	countdownDuration = 3
)

// Global variables
var (
	// The current game state
	numpadToBoard = map[ebiten.Key][2]int{
		ebiten.KeyKP1: {0, 2},
		ebiten.KeyKP2: {1, 2},
		ebiten.KeyKP3: {2, 2},
		ebiten.KeyKP4: {0, 1},
		ebiten.KeyKP5: {1, 1},
		ebiten.KeyKP6: {2, 1},
		ebiten.KeyKP7: {0, 0},
		ebiten.KeyKP8: {1, 0},
		ebiten.KeyKP9: {2, 0},
	}
)

func NewGame() *Game {
	return &Game{
		sWidth:        0,
		sHeight:       0,
		state:         StateMenu,
		gameMode:      0,
		playing:       "",
		player:        "",
		board:         [3][3]string{},
		pointsO:       0,
		pointsX:       0,
		rounds:        0,
		win:           "",
		goRythm:       nil,
		audioContext:  nil,
		audioPlayer:   nil,
		countdownTime: time.Time{},
		countdown:     countdownDuration,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.sWidth, g.sHeight
}

// Initialize the game, must be called before running the game
func (g *Game) Init(audioContext *audio.Context, sWidth, sHeight int) error {
	// Set variables
	g.sWidth = sWidth
	g.sHeight = sHeight

	// Generate the squared game board and symbols
	g.gameImage = ebiten.NewImage(sWidth, sWidth)
	g.boardImage = g.GenerateBoard(g.gameImage, sWidth)
	g.XImage, g.OImage, g.XImageHighlighted, g.OImageHighlighted, g.EmptyImage = g.GenerateSymbols(g.gameImage)

	g.randomizeStartingPlayer()

	// Initialize audio settings
	if err := g.initAudio(audioContext); err != nil {
		return err
	}

	return nil
}

// Handle game logic
func (g *Game) Update() error {
	switch g.state {

	case StateMenu:
		g.handleStateMenu()

	case StateLoading:
		err := g.handleStateLoading()
		if err != nil {
			return err
		}

	case StatePlaying:
		err := g.handleStatePlaying()
		if err != nil {
			return err
		}

	case StateGameOver:
		err := g.handleStateGameOver()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) handleStateMenu() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.state = StateLoading
		g.countdownTime = time.Now()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.gameMode = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.gameMode = 2
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.gameMode = 3
		g.goRythm = NewGoRythm()
	}
}

func (g *Game) handleStateLoading() error {
	g.player = "human"
	if g.countdown > 0 {
		elapsed := time.Since(g.countdownTime)
		if elapsed >= time.Second {
			g.countdown--
			g.countdownTime = time.Now()
		}
	} else {
		g.state = StatePlaying
		if g.audioPlayer != nil {
			g.audioPlayer.Play()
		} else {
			return fmt.Errorf("audio player is nil")
		}
	}
	return nil
}

func (g *Game) handleStatePlaying() error {
	if g.gameMode == 3 && g.goRythm.startTime.IsZero() {
		g.goRythm.Start(time.Now())
		log.LogMessage(log.DEBUG, fmt.Sprintf("Start time: %v", g.goRythm.startTime))
	}
	switch {
	// Human vs easy AI
	case g.player == "ai" && g.gameMode == 1:
		x, y := g.EasyCpu()
		g.performMove(x, y)
	// Human vs hard AI
	case g.player == "ai" && g.gameMode == 2:
		x, y := g.HardCpu()
		g.performMove(x, y)
	// Human vs human
	case g.player == "human":
		for key, pos := range numpadToBoard {
			if inpututil.IsKeyJustPressed(key) {
				x, y := pos[0], pos[1]
				if g.board[x][y] == "" {
					// GoRythm mode
					if g.gameMode == 3 {
						// Remove and highlight symbols if needed
						remove, highlight, toRemove, toHighlight := g.goRythm.Update(g.playing, x, y)
						if remove {
							g.removeSymbol(toRemove[0], toRemove[1])
						}
						if highlight {
							g.highlightSymbol(toHighlight[0], toHighlight[1])
						}
						// Calculating score on hitting the beat
						score := g.goRythm.CalculateScore()
						switch g.playing {
						case "O":
							g.pointsO += score
						case "X":
							g.pointsX += score
						}
					}
					g.performMove(x, y)
				}
			}
		}
	}
	// Check for win
	g.win, _ = g.CheckWin()
	if g.win != "" {
		g.state = StateGameOver
		if g.win == "O" {
			g.pointsO += 250
		} else {
			g.pointsX += 250
		}
	}
	// Check for draw
	if g.IsBoardFull() {
		g.state = StateGameOver
	}
	return nil
}

func (g *Game) handleStateGameOver() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// Restart the game
		g.gameImage.Clear()
		g.restartGame()

		// Return to menu and stop the music
		g.state = StateMenu
		if err := g.audioPlayer.Restart(); err != nil {
			return err
		}
	}
	return nil
}

// Restart the game and reset the variables
func (g *Game) restartGame() {
	g.board = [3][3]string{} // Reset the game board
	g.rounds = 0             // Reset the number of rounds
	g.win = ""               // Reset the win status
	g.gameMode = 0           // Reset the game mode
	g.countdown = 3          // Reset the countdown timer
	g.pointsO = 0            // Reset the points for O
	g.pointsX = 0            // Reset the points for X

	g.randomizeStartingPlayer() // Randomize the starting player
}

// Perform a move by placing the symbol, switching the player and incrementing the rounds
func (g *Game) performMove(x, y int) {
	g.placeSymbol(x, y)
	g.switchPlayer()
	g.rounds++
}

// Randomizes the starting player
func (g *Game) randomizeStartingPlayer() {
	if r := newRandom().Intn(2) == 0; r {
		g.playing = "O"
	} else {
		g.playing = "X"
	}
}

// Init the audio player with the given audio context
func (g *Game) initAudio(ctx *audio.Context) error {
	if g.audioPlayer != nil {
		g.audioPlayer.Close()
	}
	if ap, err := NewAudioPlayer(ctx); err != nil {
		log.LogMessage(log.ERROR, "failed to init audio player: "+err.Error())
		return err
	} else {
		g.audioPlayer = ap
	}
	return nil
}
