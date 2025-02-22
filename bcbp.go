package bcbp

import (
	"fmt"
	"strings"
)

const GROUP_SEPARATOR = '\x1D' // Group Separator (ASCII 0x1D)

type BCBP struct {
	Legs []*Leg `json:"legs"`
}

func (b *BCBP) String() string {

	parts := make([]string, len(b.Legs))

	for idx, l := range b.Legs {
		parts[idx] = l.String()
	}

	return strings.Join(parts, string(GROUP_SEPARATOR))
}

func Parse(raw string) (*BCBP, error) {

	parts := strings.Split(raw, string(GROUP_SEPARATOR))
	legs := make([]*Leg, len(parts))

	for idx, leg_raw := range parts {

		l, err := ParseLeg(leg_raw)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse leg at offset %d (%s), %w", idx, leg_raw, err)
		}

		legs[idx] = l
	}

	b := &BCBP{
		Legs: legs,
	}

	return b, nil
}
