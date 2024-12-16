package domain

import "context"

type DestinationProvider interface {
	SendChangelog(context.Context, string)
}
