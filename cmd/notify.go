package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/seq"
)

type NotifierReq struct {
	Data

	Title    string
	SubTitle string
	Message  string
}

type NotifierResp struct {
	Id string

	Stdout string
	Stderr string
}

func Notifier(req NotifierReq) (NotifierResp, error) {
	resp := NotifierResp{Id: seq.UUID()}

	if req.Message == "" {
		return resp, fmt.Errorf("message is empty")
	}

	args := []string{"-message", req.Message}
	if req.Title != "" {
		args = append(args, "-title", req.Title)
	}
	if req.SubTitle != "" {
		args = append(args, "-subtitle", req.SubTitle)
	}

	// args = append(args, "-sender", "com.microsoft.VSCode")

	defer tmp(req, resp)

	cmd := exec.Command("terminal-notifier", args...)

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	err := cmd.Run()

	resp.Stdout = stdout.String()
	resp.Stderr = stderr.String()

	return resp, err
}

func tmp(req NotifierReq, resp NotifierResp) {
	dir := filepath.Join(os.TempDir(), constant.NAME)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		log.Error().Msgf("mkdir error: %v", err)
	} else {
		err = os.WriteFile(filepath.Join(dir, req.FilterboxFieldPackageName, resp.Id), json.MustMarshalIndent(req), 0600)
		if err != nil {
			log.Error().Msgf("write file error: %v", err)
		}
	}
}
