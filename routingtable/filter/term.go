package filter

import (
	"github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/route"
)

// Term matches a path against a list of conditions and performs actions if it matches
type Term struct {
	from []*From
	then []Then
}

// NewTerm creates a new term
func NewTerm(from []*From, then []Then) *Term {
	t := &Term{
		from: from,
		then: then,
	}

	return t
}

// Process processes a path returning if the path should be rejected and returns a possible modified version of the path
func (t *Term) Process(p net.Prefix, pa *route.Path) (modPath *route.Path, reject bool) {
	orig := pa

	if len(t.from) == 0 {
		return t.processActions(p, pa)
	}

	for _, f := range t.from {
		if f.Matches(p, pa) {
			return t.processActions(p, pa)
		}
	}

	return orig, true
}

func (t *Term) processActions(p net.Prefix, pa *route.Path) (modPath *route.Path, reject bool) {
	modPath = pa

	for _, action := range t.then {
		modPath, reject = action.Do(p, modPath)
		if reject {
			continue
		}
	}

	return modPath, reject
}