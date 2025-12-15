package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kardianos/service"
	"github.com/farzadamr/ctf-web-tool/internal/command"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		args := strings.Split(string(body), " ")

		out := command.Run(args)
		w.Write([]byte(out))
	})

	log.Println("Engine running on 127.0.0.1:7878")
	http.ListenAndServe("127.0.0.1:7878", nil)
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	cfg := &service.Config{
		Name:        "CTFWebToolEngine",
		DisplayName: "CTF Web Tool Engine",
	}

	s, _ := service.New(&program{}, cfg)
	s.Run()
}
