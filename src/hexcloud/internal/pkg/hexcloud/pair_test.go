package hexcloud

import "testing"

type IntPair struct {
	x, y int64
}

var ip = []IntPair{
	{2, 3},
	{-2, 3},
	{0, 0},
}

var uip = []IntPair{
	{2, 3},
	{1, 2},
	{0, 0},
}

func TestSzudzikPair(t *testing.T) {
	for _, p := range uip {
		rz := SzudzikPair(p.x, p.y)
		rx, ry := SzudzikUnPair(rz)
		if rx != p.x || ry != p.y {
			t.Errorf("(Un)pairing failed, input (%d,  %d) output (%d,  %d), expected output and input being the same", p.x, p.y, rx, ry)
		}
	}
}

func TestPair(t *testing.T) {
	for _, p := range ip {
		rz := Pair(p.x, p.y)
		rx, ry := Unpair(rz)
		if rx != p.x || ry != p.y {
			t.Errorf("(Un)pairing failed, input (%d,  %d) output (%d,  %d), expected output and input being the same", p.x, p.y, rx, ry)
		}
	}
}

func TestTransforms(t *testing.T) {
	for i := -10; i < 10; i++ {
		x := SignedTransform(int64(i))
		ux := UnSignedTransform(x)
		if ux != int64(i) {
			t.Errorf("Signed and Unsigned transforms failed, input %d output %d", i, ux)
		}
	}
}
