package engine

type Entity uint32

type World struct {
	next Entity
	pos  map[Entity]*Position
	body map[Entity]*Body
}

func NewWorld() *World {
	return &World{next: 1, pos: make(map[Entity]*Position), body: make(map[Entity]*Body)}
}

func (w *World) NewEntity() Entity { e := w.next; w.next++; return e }

type Position struct{ X, Y float64 }
