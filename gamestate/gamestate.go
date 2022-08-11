package gamestate

import (
	"github.com/OlegStotsky/go-monads/maybe"
)

// TODO: comment code; use doccomments, referencing best practice; make null value meaningful for types;
// TODO: implement gesture_chain with maybe.Maybe

type Game_State struct {
	players []Player_State
	// neutral battlefield effects
	spell_defs map[string]Spell // TODO: initialize
}

type Player_State struct {
	hp              int
	gesture_history Gesture_Chain
	// controlled monsters
	// effects
}

type Spell struct {
	name           string
	incantation    Gesture_Chain
	result         []Effect
	is_enchantment bool
}

type Gesture_Chain [][2]maybe.Maybe[Gesture]

type Gesture int

const (
	nothing Gesture = iota // non-gesture
	stab                   // non-gesture

	// the gestures
	clap // both hands must clap to be valid
	fingers
	palm
	snap
	wave
	digit
)

type Effect int

const (
	// mutex debuff group
	amnesia Effect = iota
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

func New_Game_State() *Game_State {
	g := new(Game_State)

	// initialize:
	for i := 0; i < 2; i++ {
		// TODO: determine if the pointer or object is being stored here, and try to make it the pointer
		g.players = append(g.players, *New_Player_State())
	}
	// spell defs (map spell names to initialized spells)

	return g
}

func New_Player_State() *Player_State {
	p := new(Player_State)

	p.hp = 0xf

	return p
}

// String() returns a representation of the gamestate formatted for display
func (g Game_State) String() string {
	// cruft:
	// p0_left *uint64, p0_right *uint64, p1_left *uint64, p1_right *uint64

	// TODO: include more game info such as health and summons
	// FPSW DC(stab)()

	return "" // FIXME
}
