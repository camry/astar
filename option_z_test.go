package astar_test

import (
    "context"
    "testing"

    "github.com/camry/astar"
    "github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
    ctx := context.Background()
    f := astar.New(astar.Context(ctx))
    assert.Equal(t, f.Ctx(), ctx)
}

func TestMode(t *testing.T) {
    f1 := astar.New(astar.Mode(astar.Vector2Mode))
    assert.Equal(t, astar.Vector2Mode, f1.Mode())
    f2 := astar.New(astar.Mode(astar.Vector3Mode))
    assert.Equal(t, astar.Vector3Mode, f2.Mode())
}
