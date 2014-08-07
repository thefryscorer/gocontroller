package gocontroller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Button struct {
	Left, Top int
	Key       string
	Label     string
}

func (b Button) String() string {
	if b.Label != "" {
		return fmt.Sprintf(buttonTemplate, b.Left, b.Top, b.Key, b.Label)
	}
	return fmt.Sprintf(buttonTemplate, b.Left, b.Top, b.Key, b.Key)
}

type event int

const (
	NONE event = iota
	PRESS
	RELEASE
)

var eventMap = map[string]event{
	"PRESS":   PRESS,
	"RELEASE": RELEASE,
}

type input struct {
	UserIP string
	Key    string
	Event  event
}

type server struct {
	ch   chan input
	Page Layout
	Port string
}

const DefaultPort = ":12345"

func (s *server) handleRequest(w http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/" {
		io.WriteString(w, s.Page.String())
	} else {
		go s.handleInput(req)
	}

}

func (s *server) handleInput(req *http.Request) {
	if strings.Contains(req.RequestURI, "/button") {
		inputString := strings.Replace(req.RequestURI, "/button", "", 1)
		ipString := strings.Split(req.RemoteAddr, ":")[0]

		inputStrings := strings.Split(inputString, "type")

		// Check the key is allowed.
		key := inputStrings[0]
		found := false
		var foundBtn Button
		for _, btn := range s.Page.Buttons {
			if btn.Key == key {
				found = true
				foundBtn = btn
				break
			}
		}

		if found == false {
			log.Printf("Ignoring illegal key: %v", key)
			return
		}

		// If type not specified, default to release
		var ev event = RELEASE
		if len(inputStrings) > 1 {
			evString := inputStrings[1]
			ev = eventMap[evString]
		}

		in := input{
			UserIP: ipString,
			Key:    foundBtn.Key,
			Event:  ev,
		}

		s.ch <- in
	}
}

func NewServer(layout Layout, port string) *server {
	return &server{
		Port: port,
		Page: layout,
	}
}

func (s *server) Start() {
	s.ch = make(chan input)
	go func() {
		http.HandleFunc("/", s.handleRequest)
		err := http.ListenAndServe(s.Port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func (s *server) PollInput() input {
	select {
	case val := <-s.ch:
		return val
	default:
		return input{Event: NONE}
	}
}

type InputAggregator struct {
	Server *server
	Inputs []input
}

func (s *server) NewInputAggregator() InputAggregator {
	return InputAggregator{
		Inputs: make([]input, 0),
		Server: s,
	}
}

func (a *InputAggregator) Collect() {
	for i := a.Server.PollInput(); i.Event != NONE; i = a.Server.PollInput() {
		a.Inputs = append(a.Inputs, i)
	}
}

func (a *InputAggregator) Clear() {
	a.Inputs = make([]input, 0)
}
