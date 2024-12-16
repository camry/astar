package astar

import (
    "context"
)

type Option func(f *Astar)

// Context 设置上下文。
func Context(ctx context.Context) Option {
    return func(f *Astar) { f.ctx = ctx }
}

// Mode 设置向量模式。
func Mode(mode VectorMode) Option {
    return func(f *Astar) { f.mode = mode }
}
