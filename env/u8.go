package env

import (
	"strconv"
	"strings"
)

// The Saint He's specific type.
type U8 struct {
	Values []int
}

func NewU8() U8 {
	u8 := U8{
		Values: []int{},
	}
	return u8
}

func (u8 *U8) SetValues(values []int) {
	u8.Values = values
}

func (u8 *U8) String() string {
	vals := make([]string, 0, 1)
	for _, v := range u8.Values {
		vals = append(vals, strconv.Itoa(v))
	}
	return strings.Join(vals, " | ")
}

func (u8 *U8) Increment() {
	for i := range u8.Values {
		u8.Values[i]++
	}
}
