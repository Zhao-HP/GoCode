package global

import (
	"go.uber.org/zap"
	"goframe-code/pkg"
)

var Log *zap.Logger

func init() {
	Log = pkg.Zap()
}
