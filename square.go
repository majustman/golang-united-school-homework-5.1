package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq *Square) End() Point {
	// implement me
	return Point{
		x: sq.start.x + int(sq.a),
		y: sq.start.y - int(sq.a),
	}
}

func (sq *Square) Area() uint {
	// implement me
	return sq.a * sq.a
}

func (sq *Square) Perimeter() uint {
	// implement me
	return sq.a * uint(4)
}
