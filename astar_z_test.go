package astar_test

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/astar"
)

func TestNew(t *testing.T) {
    tiles := make(map[string]*astar.Tile, 49)
    start := 0
    for y := 0; y < 7; y++ {
        if y > 0 && y%2 == 0 {
            start = -(y / 2)
        }
        for x := start; x < start+7; x++ {
            tiles[fmt.Sprintf("%d:%d", x, y)] = astar.NewTile(astar.NewVector(int32(x), int32(y)))
            // fmt.Printf("%2d,%2d,%2d  ", x, y, 0-y-x)
        }
        // fmt.Println()
    }
    obstacles := map[string]*astar.Vector{
        "2:2":  astar.NewVector(2, 2),
        "2:3":  astar.NewVector(2, 3),
        "-1:4": astar.NewVector(-1, 4),
        "1:4":  astar.NewVector(1, 4),
    }
    for s := range obstacles {
        if tile, ok := tiles[s]; ok {
            tile.SetIsObstacle(true)
        }
    }
    for _, tile := range tiles {
        // 左
        leftK := fmt.Sprintf("%d:%d", tile.Vector().X()-1, tile.Vector().Y())
        if leftTile, ok := tiles[leftK]; ok {
            tile.AddNearList(leftTile)
        }
        // 右
        rightK := fmt.Sprintf("%d:%d", tile.Vector().X()+1, tile.Vector().Y())
        if rightTile, ok := tiles[rightK]; ok {
            tile.AddNearList(rightTile)
        }
        // 左上
        leftUpK := fmt.Sprintf("%d:%d", tile.Vector().X()-1, tile.Vector().Y()+1)
        if leftUpTile, ok := tiles[leftUpK]; ok {
            tile.AddNearList(leftUpTile)
        }
        // 右下
        rightDownK := fmt.Sprintf("%d:%d", tile.Vector().X()+1, tile.Vector().Y()-1)
        if rightDownTile, ok := tiles[rightDownK]; ok {
            tile.AddNearList(rightDownTile)
        }
        // 左下
        leftDownK := fmt.Sprintf("%d:%d", tile.Vector().X(), tile.Vector().Y()-1)
        if leftDownTile, ok := tiles[leftDownK]; ok {
            tile.AddNearList(leftDownTile)
        }
        // 右上
        rightUpK := fmt.Sprintf("%d:%d", tile.Vector().X(), tile.Vector().Y()+1)
        if rightUpTile, ok := tiles[rightUpK]; ok {
            tile.AddNearList(rightUpTile)
        }
    }
    startTile := tiles["0:3"]
    endTile := tiles["3:4"]
    pathVector := []*astar.Vector{
        astar.NewVector(0, 3),
        astar.NewVector(0, 4),
        astar.NewVector(0, 5),
        astar.NewVector(1, 5),
        astar.NewVector(2, 5),
        astar.NewVector(3, 4),
    }
    paths := astar.New(astar.Mode(astar.Vector3Mode)).FindPath(startTile, endTile)
    for i, path := range paths {
        assert.Equal(t, path.Vector().X(), pathVector[i].X())
        assert.Equal(t, path.Vector().Y(), pathVector[i].Y())
    }
}
