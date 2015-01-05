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
	Color     string
	Style     string
}

func (b Button) String() string {
	var style string
	var col string
	if b.Color == "" {
		col = ""
	} else {
		col = fmt.Sprintf("background:%v;", b.Color)
	}
	style = b.Style + col
	if b.Label != "" {
		return fmt.Sprintf(buttonTemplate, b.Left, b.Top, style, b.Key, b.Label)
	}
	return fmt.Sprintf(buttonTemplate, b.Left, b.Top, style, b.Key, b.Key)
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

type Input struct {
	UserIP string
	Key    string
	Event  event
}

type Server struct {
	ch   chan Input
	Page Layout
	Port string
}

const DefaultPort = ":12345"

func (s *Server) handleRequest(w http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/" {
		io.WriteString(w, s.Page.String())
	} else {
		go s.handleInput(req)
	}

}

func (s *Server) handleInput(req *http.Request) {
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

		in := Input{
			UserIP: ipString,
			Key:    foundBtn.Key,
			Event:  ev,
		}

		s.ch <- in
	}
}

func NewServer(layout Layout, port string) *Server {
	return &Server{
		Port: port,
		Page: layout,
	}
}

func (s *Server) Start() {
	s.ch = make(chan Input)
	go func() {
		http.HandleFunc("/", s.handleRequest)
		err := http.ListenAndServe(s.Port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func (s *Server) PollInput() Input {
	select {
	case val := <-s.ch:
		return val
	default:
		return Input{Event: NONE}
	}
}

type InputAggregator struct {
	Server *Server
	Inputs []Input
}

func (s *Server) NewInputAggregator() InputAggregator {
	return InputAggregator{
		Inputs: make([]Input, 0),
		Server: s,
	}
}

func (a *InputAggregator) Collect() {
	for i := a.Server.PollInput(); i.Event != NONE; i = a.Server.PollInput() {
		a.Inputs = append(a.Inputs, i)
	}
}

func (a *InputAggregator) Clear() {
	a.Inputs = make([]Input, 0)
}
