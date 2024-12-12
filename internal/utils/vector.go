package utils

import (
	"math"

	"github.com/charmbracelet/log"
)

type IVec3 [3]int64

// Deprecated: Use IVec instead
type IVec2 [2]int64
type FVec [3]float64

type IVec struct {
	X, Y int
}

func (ivec *IVec) Add(add IVec) IVec {
	return IVec{
		X: ivec.X + add.X,
		Y: ivec.Y + add.Y,
	}
}
func (ivec *IVec) Sub(add IVec) IVec {
	return IVec{
		X: ivec.X - add.X,
		Y: ivec.Y - add.Y,
	}
}
func (ivec *IVec) Rotate(degrees int) IVec {
	radians := float64(degrees) * math.Pi / 180

	return IVec{
		X: int(float64(ivec.X)*math.Cos(radians) - float64(ivec.Y)*math.Sin(radians)),
		Y: int(float64(ivec.X)*math.Sin(radians) + float64(ivec.Y)*math.Cos(radians)),
	}
}
func (ivec *IVec) Scale(scale int) IVec {
	return IVec{
		X: ivec.X * scale,
		Y: ivec.Y * scale,
	}
}

func (ivec *IVec2) Add(add IVec2) {
	if len(ivec) != 2 {
		log.Fatal("IVec2.Add(): vector did not contain 2 values, needed to increment", "actual", ivec)
	}
	if len(add) != 2 {
		log.Fatal("IVec2.Add(): parameter did not contain 2 values, needed to increment", "actual", add)
	}
	ivec[0] += add[0]
	ivec[1] += add[1]
}
func (ivec *IVec2) Sub(sub IVec2) {
	if len(ivec) != 2 {
		log.Fatal("IVec2.Sub(): vector did not contain 2 values, needed to decrement", "actual", ivec)
	}
	if len(sub) != 2 {
		log.Fatal("IVec2.Sub(): parameter did not contain 2 values, needed to decrement", "actual", sub)
	}
	ivec[0] -= sub[0]
	ivec[1] -= sub[1]
}
func (ivec *IVec2) Rotate(degrees int) {
	if len(ivec) != 2 {
		log.Fatal("IVec2.Rotate(): vector did not contain 2 values, needed to rotate", "actual", ivec)
	}
	radians := float64(degrees) * math.Pi / 180

	ivec[0] = int64(float64(ivec[0])*math.Cos(radians) - float64(ivec[1])*math.Sin(radians))
	ivec[1] = int64(float64(ivec[0])*math.Sin(radians) + float64(ivec[1])*math.Cos(radians))
}
func (ivec *IVec2) Scale(scale int) {
	if len(ivec) != 2 {
		log.Fatal("IVec2.Scale(): vector did not contain 2 values, needed to scale", "actual", ivec)
	}
	ivec[0] *= int64(scale)
	ivec[1] *= int64(scale)
}
func (ivec *IVec2) Equals(vec IVec2) bool {
	if len(ivec) != 2 {
		log.Fatal("IVec2.Equals(): vector did not contain 2 values, needed to compare", "actual", ivec)
	}
	if len(vec) != 2 {
		log.Fatal("IVec2.Equals(): parameter did not contain 2 values, needed to compare", "actual", vec)
	}
	return ivec[0] == vec[0] && ivec[1] == vec[1]
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (direction *Direction) Rotate() Direction {
	switch *direction {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	default:
		return North
	}
}
func ToIVec(direction Direction) IVec {
	switch direction {
	case North:
		return IVec{X: 0, Y: -1}
	case East:
		return IVec{X: 1, Y: 0}
	case South:
		return IVec{X: 0, Y: 1}
	default:
		return IVec{X: -1, Y: 0}
	}
}
func ToIVec2(direction Direction) IVec2 {
	switch direction {
	case North:
		return IVec2{0, -1}
	case East:
		return IVec2{1, 0}
	case South:
		return IVec2{0, 1}
	default:
		return IVec2{-1, 0}
	}
}
