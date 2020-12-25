package main

import "fmt"

type Move int

const (
	UNKNOWN Move = 0
	W       Move = 1
	E       Move = 2
	N       Move = 10
	S       Move = 20
	NW      Move = 11
	NE      Move = 12
	SW      Move = 21
	SE      Move = 22
)

func (m Move) String() string {
	switch m {
	case UNKNOWN:
		return "UNKNOWN"
	case W:
		return "W"
	case E:
		return "E"
	case N:
		return "N"
	case S:
		return "S"
	case NW:
		return "NW"
	case NE:
		return "NE"
	case SW:
		return "SW"
	case SE:
		return "SE"
	default:
		return fmt.Sprintf("OTHER(%d)", m)
	}
}

func MoveFromChar(c uint8) (Move, error) {
	switch c {
	case 'w':
		return W, nil
	case 'e':
		return E, nil
	case 'n':
		return N, nil
	case 's':
		return S, nil
	default:
		return 0, fmt.Errorf("cannot convert '%c' to move", c)
	}
}

func ParsePath(input string) ([]Move, error) {
	path := make([]Move, 0, len(input))

	state := UNKNOWN
	for i := 0; i < len(input); i++ {
		currentMove, err := MoveFromChar(input[i])
		if err != nil {
			return nil, fmt.Errorf("unexpected state (state=%s, index=%d): %w", state, i, err)
		}

		if currentMove == W || currentMove == E {
			path = append(path, state+currentMove)
			state = UNKNOWN
			continue
		}

		if state == UNKNOWN && (currentMove == N || currentMove == S) {
			state = currentMove
			continue
		}

		return nil, fmt.Errorf("unexpected state (current=%s, state=%s, index=%d)", currentMove, state, i)
	}

	if state != UNKNOWN {
		return nil, fmt.Errorf("unexpected end of input (state=%s)", state)
	}

	return path, nil
}

func ParsePaths(input []string) ([][]Move, error) {
	paths := make([][]Move, len(input))
	for i, l := range input {
		path, err := ParsePath(l)
		if err != nil {
			return nil, err
		}
		paths[i] = path
	}
	return paths, nil
}
