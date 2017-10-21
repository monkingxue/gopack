package frame

type (
	Scope struct {
		parent       []*Scope
		depth        int
		variables    []string
		isBlockScope bool
	}
)
