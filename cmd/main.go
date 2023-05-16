package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os/exec"
	"strings"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/log"
)

func main() {
	app.Add(start)
	app.Defer(stop)
	err := app.Go()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
}

var server *http.Server

func start(context.Context) error {
	http.HandleFunc("/", index)

	addr := ":5400"

	server = &http.Server{Handler: http.DefaultServeMux}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Info().Msgf("api start success on %s", addr)

	return server.Serve(ln)
}

func stop() {
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Error().Msgf("api shutdown error: %v", err)
	}
}

type Data struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	App   string `json:"app"`
}

func index(_ http.ResponseWriter, r *http.Request) {
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Msgf("read body error: %v", err)
		return
	}

	log.Info().Msgf("msg: %s", bs)

	data, err := json.UnmarshalTo[*Data](bs)
	if err != nil {
		log.Error().Msgf("decode error: %v", err)
		return
	}

	if data.App == "" || data.Text == "" {
		return
	}

	notify(fmt.Sprintf("%s(%s)", data.Title, data.App), data.Text)
}

func notify(title, content string) {
	stdout, stderr, err := run(fmt.Sprintf(`display notification "%s" with title "%s"`, content, title))
	if err != nil {
		panic(err)
	}
	log.Info().Msgf("stdout: {%s}, stderr: {%s}, err: {%v}", stdout, stderr, err)
}

func run(code string) (string, string, error) {
	cmd := exec.Command("osascript", "-l", "AppleScript")
	cmd.Stdin = strings.NewReader(code)

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
