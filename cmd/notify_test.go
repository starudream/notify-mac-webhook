package main

import (
	"testing"

	"github.com/starudream/go-lib/testx"
)

func TestNotifier(t *testing.T) {
	resp, err := Notifier(NotifierReq{
		Title:   "This is notification",
		Message: `lorem ipsum dolor sit amet, consectetur adipiscing elit.`,
	})
	testx.P(t, err, resp)
}
