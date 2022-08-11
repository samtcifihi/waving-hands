package gamestate

import (
	"github.com/OlegStotsky/go-monads/maybe"
)

// TODO: comment code; use doccomments, referencing best practice; make null value meaningful for types;

type game_state struct {
	players []player_state
	// neutral battlefield effects
	spell_defs map[string]spell // TODO: initialize
}

type player_state struct {
	hp              int
	gesture_history gesture_chain
	// controlled monsters
	// effects
}

type spell struct {
	name           string
	incantation    gesture_chain
	result         []effect
	is_enchantment bool
}

type gesture_chain [][2]maybe.Maybe[gesture]

type gesture int

const (
	fingers gesture = iota
	palm
	snap
	wave
	digit
	clap
	stab
	nothing
)

type effect int

const (
	// mutex debuff group
	amnesia effect = iota
	confusion
	charm_person
	charm_monster
	paralysis
	fear

	// dot debuffs
	disease
	poison

	// action debuff
	blindness

	// action buff
	haste

	// resistances
	protection_from_evil
	resist_heat
	resist_cold
	invisibility
)

func New_Game_State() *game_state {
	g := new(game_state)

	// initialize:
	for i := 0; i < 2; i++ {
		// TODO: determine if the pointer or object is being stored here, and try to make it the pointer
		g.players = append(g.players, *New_Player_State())
	}
	// spell defs (map spell names to initialized spells)

	return g
}

func New_Player_State() *player_state {
	p := new(player_state)

	p.hp = 0xf

	return p
}

// String() returns a representation of the gamestate formatted for display
func (g game_state) String() string {
	// cruft:
	// p0_left *uint64, p0_right *uint64, p1_left *uint64, p1_right *uint64

	// TODO: include more game info such as health and summons
	// FPSW DC(stab)()

	return "" // FIXME
}
