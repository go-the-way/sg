package orderby

import (
	"github.com/billcoding/sgen"
)

var (
	Gen       = func(gs ...sgen.Ge) sgen.Ge { return sgen.NewJoiner(gs, ", ", " ORDER BY ", "", false) }
	Asc       = func(g sgen.Ge) sgen.Ge { return sgen.NewJoiner([]sgen.Ge{g}, "", "", " ASC", false) }
	Desc      = func(g sgen.Ge) sgen.Ge { return sgen.NewJoiner([]sgen.Ge{g}, "", "", " DESC", false) }
	AscGroup  = func(gs ...sgen.Ge) sgen.Ge { return sgen.NewJoinerWithAppend(gs, ", ", "", "", " ASC", false) }
	DescGroup = func(gs ...sgen.Ge) sgen.Ge { return sgen.NewJoinerWithAppend(gs, ", ", "", "", " DESC", false) }
)
