package workflow

import (
	"github.com/bitmagnet-io/bitmagnet/internal/classifier"
	"github.com/bitmagnet-io/bitmagnet/internal/processor/parsers"
)

const parseVideoContentName = "parse_video_content"

type parseVideoContentAction struct{}

func (parseVideoContentAction) Name() string {
	return parseVideoContentName
}

var parseVideoContentPayloadSpec = payloadLiteral[string]{parseVideoContentName}

func (parseVideoContentAction) compileAction(ctx compilerContext) (action, error) {
	if _, err := parseVideoContentPayloadSpec.Unmarshal(ctx); err != nil {
		return action{}, ctx.error(err)
	}
	return action{
		run: func(ctx executionContext) (classifier.Classification, error) {
			parsed, err := parsers.ParseVideoContent(ctx.result.ContentType, ctx.torrent.Name)
			cl := ctx.result
			if err == nil {
				cl.Merge(parsed)
			}
			return cl, nil
		},
	}, nil
}
