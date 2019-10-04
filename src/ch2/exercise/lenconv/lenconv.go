package lenconv

import "fmt"

type CM float64
type DM float64
type M float64

const (
	OneM M = 1
	TenM  M = 10
)

func (m M) String() string {
	return fmt.Sprintf("%g m", m)
}

func (dm DM) String() string {
	return fmt.Sprintf("%g dm", dm)
}

func (cm CM) String() string {
	return fmt.Sprintf("%g cm", cm)
}
