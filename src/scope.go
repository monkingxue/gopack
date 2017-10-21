package src

type (
	Scope struct {
		Parent  *Scope
		Depth   int
		Vars    []string
		IsBlock bool
	}
)

func CreateScope(p *Scope, n string, ib bool) Scope {
	var depth int
	if p != nil {
		depth = p.Depth + 1
	}
	return Scope{Parent: p, Depth: depth, Vars: n, IsBlock: ib}
}

func (s *Scope) Add(name string, hoist bool) {
	if !hoist && s.IsBlock {
		s.Parent.Add(name, hoist)
	} else {
		s.Vars = append(s.Vars, name)
	}
}

func (s *Scope) findScope(name string) Scope {
	for _, v := range s.Vars {
		if v == name {
			return *s
		}
	}

	if s.Parent != nil {
		return s.Parent.findScope(name)
	}

	return Scope{Depth: -1}
}

func (s *Scope) contains(name string) bool {
	return s.findScope(name).Depth != -1
}
