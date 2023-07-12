package main

import (
	"testing"

	"github.com/starudream/go-lib/seq"
	"github.com/starudream/go-lib/testx"
)

func TestNotify(t *testing.T) {
	resp, err := Notify(seq.UUID(), NotifyReq{
		Title:   "This is notification",
		Message: `lorem ipsum dolor sit amet, consectetur adipiscing elit.`,
	})
	testx.P(t, err, resp)
}
