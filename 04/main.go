package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Grid[T any] struct {
	numCols, numRows int
	cells            []T
}

func NewGrid[T any](numCols, numRows int) *Grid[T] {
	return &Grid[T]{
		numRows: numRows,
		numCols: numCols,
		cells:   make([]T, numRows*numCols),
	}
}

func (g Grid[T]) indexOf(x, y int) int {
	return y*g.numCols + x
}

func (g Grid[T]) Get(x, y int) T {
	return g.cells[g.indexOf(x, y)]
}

func (g *Grid[T]) Set(x, y int, val T) {
	g.cells[g.indexOf(x, y)] = val
}

func (g Grid[T]) InBounds(x, y int) bool {
	return x >= 0 && x < g.numCols && y >= 0 && y < g.numRows
}

func (g Grid[T]) NumRows() int { return g.numRows }

func (g Grid[T]) NumCols() int { return g.numCols }

func (g *Grid[T]) AddRow() {
	g.cells = append(g.cells, make([]T, g.numCols)...)
	g.numRows++
}

func (g Grid[T]) ConditionalCount(fn func(T) bool) int {
	ret := 0
	for _, v := range g.cells {
		if fn(v) {
			ret++
		}
	}

	return ret
}

func (g *Grid[T]) Print() {
	for y := range g.numRows {
		for x := range g.numCols {
			fmt.Print(g.Get(x, y))
		}

		fmt.Println()
	}
}

type Vec2 struct {
	X, Y int
}

func main() {
	fmt.Println(partTwo())
}

func partOne() int {
	return buildAdjacencyGrid().ConditionalCount(func(i int) bool { return i < 4 })
}

const EMPTY = 9

func partTwo() int {
	g := buildAdjacencyGrid()
	ret := new(int)
	countAndPruneAdjacent(g, ret)
	return *ret
}

var offsets8 = []Vec2{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}

func countAndPruneAdjacent(g *Grid[int], acc *int) {
	toAdd := 0
	for y := range g.NumRows() {
		for x := range g.NumCols() {
			adj := g.Get(x, y)
			if adj >= 4 {
				continue
			}

			toAdd += 1
			g.Set(x, y, EMPTY)

			for _, offset := range offsets8 {
				offX, offY := x+offset.X, y+offset.Y
				if !g.InBounds(offX, offY) {
					continue
				}

				offAdj := g.Get(offX, offY)
				if offAdj >= EMPTY {
					continue
				}

				g.Set(offX, offY, offAdj-1)
			}
		}
	}

	if toAdd == 0 {
		return
	}

	*acc += toAdd
	countAndPruneAdjacent(g, acc)
}

func buildAdjacencyGrid() *Grid[int] {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	s := string(b)

	firstLine := s[:strings.IndexByte(s, '\n')]
	numCols := len(firstLine)

	g := NewGrid[int](numCols, 0)

	offsets := []Vec2{{-1, 0}, {-1, -1}, {0, -1}, {1, -1} /*, {1, 0}, {1, 1}, {0, 1}, {-1, 1}*/}

	for line := range strings.Lines(s) {
		if len(line) == 0 {
			continue
		}

		y := g.NumRows()
		g.AddRow()

		for x := range numCols {
			if line[x] != '@' {
				g.Set(x, y, EMPTY)
				continue
			}

			for _, offset := range offsets {
				offX, offY := x+offset.X, y+offset.Y

				if g.InBounds(offX, offY) {
					curAdjToOffset := g.Get(offX, offY)

					if curAdjToOffset >= EMPTY {
						continue
					}

					g.Set(offX, offY, curAdjToOffset+1)
					curAdj := g.Get(x, y)
					g.Set(x, y, curAdj+1)
				}
			}
		}
	}

	return g
}
