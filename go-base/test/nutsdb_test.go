package test

import (
	"github.com/xujiajun/nutsdb"
	"testing"
)

func TestOpen(t *testing.T) {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/Users/zhao/Desktop/Application/nutsdb"
	_, err := nutsdb.Open(opt)
	if err == nil {
		return
	}

}
