package main

type Game_State struct {
	players []player_state
	// neutral battlefield effects
	spell_defs map[string]spell // TODO: initialize
}

type player_state struct {
	hp int
	l_hand gesture_chain
	r_hand gesture_chain
	// controlled monsters
	// effects
}

type spell struct {
	name string
	incantation gesture_chain
	result []effect
	is_enchantment bool
}

type gesture_chain struct {
	incantation []gesture
	is_both_hands []bool
}

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

// Format() returns a representation of the gamestate formatted for display
func (g Game_State) Format() string {
	// cruft:
	// p0_left *uint64, p0_right *uint64, p1_left *uint64, p1_right *uint64

	// TODO: include more game info such as health and summons
	// FPSW DC(stab)()

	return "" // FIXME
}
