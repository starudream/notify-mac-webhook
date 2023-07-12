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
)

type NotifyReq struct {
	Data

	Title    string
	SubTitle string
	Message  string
}

type NotifyResp struct {
	Id string

	Stdout string
	Stderr string
}

func Notify(id string, req NotifyReq) (resp NotifyResp, err error) {
	resp = NotifyResp{Id: id}

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

	args = append(args, "-timeout", "30")
	args = append(args, "-actions", "Close")

	defer tmp(req, resp)

	cmd := exec.Command("alerter", args...)

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	err = cmd.Run()

	resp.Stdout = stdout.String()
	resp.Stderr = stderr.String()

	return resp, err
}

func tmp(req NotifyReq, resp NotifyResp) {
	dir := filepath.Join(os.TempDir(), constant.NAME)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		log.Error().Msgf("mkdir error: %v", err)
	} else {
		err = os.WriteFile(filepath.Join(dir, req.FilterboxFieldPackageName+"-"+resp.Id), json.MustMarshalIndent(map[string]any{"req": req, "resp": resp}), 0600)
		if err != nil {
			log.Error().Msgf("write file error: %v", err)
		}
	}
}
