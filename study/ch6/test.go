package ch6

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (Path Point) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point // 方法变量
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i, _ := range path {
		// 调用 path[i].Add(offset) 或者是 path[i].Sub(offset)
		path[i] = op(path[i], offset) // 等价于 path[i].Sub/Add(offset)
	}
}
