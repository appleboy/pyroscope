package tree

import "github.com/appleboy/pyroscope/pkg/structs/cappedarr"

func (t *Tree) minValue(maxNodes int) uint64 {
	c := cappedarr.New(maxNodes)
	t.iterateWithCum(func(cum uint64) bool {
		return c.Push(cum)
	})
	return c.MinValue()
}
