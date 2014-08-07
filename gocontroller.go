package gocontroller

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type Button int

const (
	NONE Button = iota
	UP
	DOWN
	LEFT
	RIGHT
	B
	A
	START
	SELECT
)

type inputType int

const (
	PRESS inputType = iota
	RELEASE
)

var typeMap = map[string]inputType{
	"PRESS":   PRESS,
	"RELEASE": RELEASE,
}

type input struct {
	UserIP    string
	Button    Button
	InputType inputType
}

var inputMap = map[string]Button{
	"UP":     UP,
	"DOWN":   DOWN,
	"LEFT":   LEFT,
	"RIGHT":  RIGHT,
	"START":  START,
	"SELECT": SELECT,
	"A":      A,
	"B":      B,
}

type server struct {
	ch   chan input
	Page string // For now
	Port string
}

const DEFAULTPAGE = gamepadPage
const DEFAULTPORT = ":12345"

func (s *server) handleRequest(w http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/" {
		io.WriteString(w, s.Page)
	} else {
		s.handleInput(req)
	}

}

func (s *server) handleInput(req *http.Request) {
	if strings.Contains(req.RequestURI, "/button") {
		inputString := strings.Replace(req.RequestURI, "/button", "", 1)
		ipString := strings.Split(req.RemoteAddr, ":")[0]

		inputStrings := strings.Split(inputString, "type")

		buttonString := inputStrings[0]
		button := inputMap[buttonString]

		// If type not specified, default to release
		var inType inputType = RELEASE
		if len(inputStrings) > 1 {
			typeString := inputStrings[1]
			inType = typeMap[typeString]
		}

		event := input{
			UserIP:    ipString,
			Button:    button,
			InputType: inType,
		}

		s.ch <- event
	}
}

func NewServer(page string, port string) *server {
	return &server{
		Port: port,
		Page: page,
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
		return input{
			UserIP: "",
			Button: NONE,
		}
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
	for i := a.Server.PollInput(); i.Button != NONE; i = a.Server.PollInput() {
		a.Inputs = append(a.Inputs, i)
	}
}

func (a *InputAggregator) Clear() {
	a.Inputs = make([]input, 0)
}
