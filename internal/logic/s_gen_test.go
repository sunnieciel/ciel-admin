package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestTableFields(t *testing.T) {
	g.Dump(Gen.TableFields(context.TODO(), "u_user"))
}
