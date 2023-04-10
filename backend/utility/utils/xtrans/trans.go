package xtrans

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func T(language interface{}, word string) string {
	if language == nil {
		language = "en"
	}
	var (
		ctx  = context.TODO()
		i18n = g.I18n()
	)
	if err := i18n.SetPath(fmt.Sprint("/resource/i18n")); err != nil {
		panic(err)
	}
	i18n.SetLanguage(fmt.Sprint(language))
	return i18n.Translate(ctx, word)
}
