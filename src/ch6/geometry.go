package main

import "math"

import "image/color"

type Point struct {
	X, Y float64
}

type Path []Point

type ColorPoint struct {
	Point
	color.RGBA
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, p.Y-q.Y)
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, p.Y-q.Y)
}

func (path Path) Distance() float64 {
	total := 0.0
	for i := range path {
		if i > 0 {
			total += path[i-1].Distance(path[i])
		}
	}
	return total
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (path Path) TranslateBy(offset Point, op func(Point, Point) Point) Path {
	for i := range path {
		path[i] = op(path[i], offset)
	}
	return path
}
