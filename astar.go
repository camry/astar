package astar

import (
    "context"
    "reflect"
    "sort"

    "github.com/camry/fp"
    "github.com/samber/lo"
)

type VectorMode int8

const (
    Vector2Mode VectorMode = iota + 1 // 二维模式
    Vector3Mode                       // 三维模式
)

type Astar struct {
    ctx  context.Context // 上下文。
    mode VectorMode      // 向量模式。
}

// New 新建寻路对象。
func New(opts ...Option) *Astar {
    f := &Astar{
        ctx:  context.Background(),
        mode: Vector2Mode,
    }
    for _, opt := range opts {
        opt(f)
    }
    return f
}

// Ctx 上下文。
func (a *Astar) Ctx() context.Context {
    return a.ctx
}

// Mode 向量模式。
func (a *Astar) Mode() VectorMode {
    return a.mode
}

// FindPath 寻找路径。
func (a *Astar) FindPath(start, end *Tile) []*Tile {
    openPathTiles := make([]*Tile, 0, 50)
    closedPathTiles := make([]*Tile, 0, 50)
    finalPathTiles := make([]*Tile, 0, 50)

    currentTile := start
    currentTile.g = 0
    if a.Mode() == Vector3Mode {
        currentTile.h = a.Vector3H(start.vector, end.vector)
    } else {
        currentTile.h = a.Vector2H(start.vector, end.vector)
    }

    openPathTiles = append(openPathTiles, currentTile)

    for len(openPathTiles) > 0 {
        // 对打开的列表进行排序，以获得 F 值最低的那块地格。
        sort.SliceStable(openPathTiles, func(i, j int) bool {
            if openPathTiles[i].F() < openPathTiles[j].F() {
                return true
            }
            if openPathTiles[i].g > openPathTiles[j].g {
                return true
            }
            return false
        })
        currentTile = openPathTiles[0]

        // 将当前地格从开放列表中移除，并将其添加到封闭列表中。
        for index, tile := range openPathTiles {
            if reflect.DeepEqual(tile, currentTile) {
                openPathTiles = append(openPathTiles[:index], openPathTiles[index+1:]...)
            }
        }
        closedPathTiles = append(closedPathTiles, currentTile)

        g := currentTile.g + 1

        // 如果在关闭的列表中有一个目标地格，我们就找到了一个路径。
        if lo.Contains(closedPathTiles, end) {
            break
        }

        // 调查当前地格的每一块相邻的地格。
        for _, tile := range currentTile.nearList {
            if tile.isObstacle {
                continue
            }
            if lo.Contains(closedPathTiles, tile) {
                continue
            }
            // 如果它不在开放列表中，则添加它并计算 G 和 H。
            if !(lo.Contains(openPathTiles, tile)) {
                tile.g = g
                if a.Mode() == Vector3Mode {
                    tile.h = a.Vector3H(tile.vector, end.vector)
                } else {
                    tile.h = a.Vector2H(tile.vector, end.vector)
                }
                openPathTiles = append(openPathTiles, tile)
            } else if tile.F() > g+tile.h {
                tile.g = g
            }
        }
    }

    // 回溯设置最终路径。
    if lo.Contains(closedPathTiles, end) {
        currentTile = end
        finalPathTiles = append(finalPathTiles, currentTile)
        for i := end.g - 1; i >= 0; i-- {
            for _, tile := range closedPathTiles {
                if tile.g == i && lo.Contains(currentTile.nearList, tile) {
                    currentTile = tile
                }
            }
            finalPathTiles = append(finalPathTiles, currentTile)
        }
        sort.SliceStable(finalPathTiles, func(i, j int) bool {
            return true
        })
    }

    return finalPathTiles
}

// Vector2H H = |x1 – x2| + |y1 – y2|。
func (a *Astar) Vector2H(start, end *Vector) int32 {
    d1 := fp.F64FromInt32(start.X()).Sub(fp.F64FromInt32(end.X())).Abs()
    d2 := fp.F64FromInt32(start.Y()).Sub(fp.F64FromInt32(end.Y())).Abs()
    return d1.Add(d2).FloorToInt()
}

// Vector3H H = Max(|x1 – x2|, |y1 – y2|, |z1 – z2|)。
func (a *Astar) Vector3H(start, end *Vector) int32 {
    d1 := fp.F64FromInt32(start.X()).Sub(fp.F64FromInt32(end.X())).Abs()
    d2 := fp.F64FromInt32(start.Y()).Sub(fp.F64FromInt32(end.Y())).Abs()
    d3 := fp.F64FromInt32(start.Z()).Sub(fp.F64FromInt32(end.Z())).Abs()
    return fp.F64Max(d1, d2, d3).FloorToInt()
}
