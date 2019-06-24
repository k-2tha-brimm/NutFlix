package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/k-2tha-brimm/clone/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
