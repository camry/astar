package astar

// Tile 地格。
type Tile struct {
    f          int32   // f 是 g 和 h 的总和。
    g          int32   // 从起点地格到此地格的成本。
    h          int32   // 从这个地格到目标地格估计成本。
    vector     *Vector // 地格坐标。
    nearList   []*Tile // 相邻地格列表。
    isObstacle bool    // 是否阻挡？
}

// NewTile 新建地格。
func NewTile(vector *Vector) *Tile {
    return &Tile{
        vector: vector,
    }
}

// F 是 g 和 h 的总和。
func (t *Tile) F() int32 {
    return t.g + t.h
}

// Vector 地格坐标。
func (t *Tile) Vector() *Vector {
    return t.vector
}

// NearList 相邻地格列表。
func (t *Tile) NearList() []*Tile {
    return t.nearList
}

// AddNearList 添加相邻地格。
func (t *Tile) AddNearList(tile *Tile) {
    t.nearList = append(t.nearList, tile)
}

// IsObstacle 是否阻挡？
func (t *Tile) IsObstacle() bool {
    return t.isObstacle
}

// SetIsObstacle 设置是否阻挡。
func (t *Tile) SetIsObstacle(isObstacle bool) {
    t.isObstacle = isObstacle
}
