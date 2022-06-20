package main

func new_bitmap(val uint64) *uint64 {
	/*
	covers one hand's most recent 8 gestures:
		     .
		     .
		     ...
		     FPSW DC(stab)()
		     89ab cde     f
	most recent: FPSW DC(stab)()
		     0123 456     7
	*/
	x := val
	return &x
}

type Bitmaps struct {
	empty	uint64
	full	uint64
	fingers	uint64
	palm	uint64
	snap	uint64
	wave	uint64
	digit	uint64
	clap	uint64
	stab	uint64
	nothing	uint64
}

func default_bitmaps() *Bitmaps {
	b := new(Bitmaps)
	b.empty = 0
	b.full = 0
	b.full = b.full - 1
	b.fingers = 0x8080808080808080
	b.palm = 0x4040404040404040

	return b
}
