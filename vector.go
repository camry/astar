package astar

// Vector 坐标。
type Vector struct {
    x int32
    y int32
}

// NewVector 新建坐标。
func NewVector(x, y int32) *Vector {
    return &Vector{x: x, y: y}
}

func (v *Vector) X() int32 {
    return v.x
}

func (v *Vector) Y() int32 {
    return v.y
}

func (v *Vector) Z() int32 {
    return 0 - v.x - v.y
}
