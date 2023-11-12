package dashlinux

import "context"

type BuildInfo struct {
	Version string
	Commit  string
}

type BuildStore interface {
	Get(context.Context) (BuildInfo, error)
	Update(context.Context, BuildInfo) error
}
