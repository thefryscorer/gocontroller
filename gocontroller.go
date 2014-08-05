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
)

type input struct {
	UserIP string
	Button Button
}

var ch chan input

func showIndex(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, gamepadPage)
}

func StartServer() {
	ch = make(chan input)
	go func() {
		http.HandleFunc("/", showIndex)
		http.HandleFunc("/buttonUP", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: UP,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		http.HandleFunc("/buttonDOWN", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: DOWN,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		http.HandleFunc("/buttonLEFT", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: LEFT,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		http.HandleFunc("/buttonRIGHT", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: RIGHT,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		http.HandleFunc("/buttonB", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: B,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		http.HandleFunc("/buttonA", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: A,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		http.HandleFunc("/buttonSTART", func(w http.ResponseWriter, req *http.Request) {
			ch <- input{
				Button: START,
				UserIP: strings.Split(req.RemoteAddr, ":")[0],
			}
		})
		err := http.ListenAndServe(":12345", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func PollInput() input {
	select {
	case val := <-ch:
		return val
	default:
		return input{
			UserIP: "",
			Button: NONE,
		}
	}
}

type InputAggregator struct {
	Inputs []input
}

func NewInputAggregator() InputAggregator {
	return InputAggregator{
		Inputs: make([]input, 0),
	}
}

func (a *InputAggregator) Collect() {
	for i := PollInput(); i.Button != NONE; i = PollInput() {
		a.Inputs = append(a.Inputs, i)
	}
}

func (a *InputAggregator) Clear() {
	a.Inputs = make([]input, 0)
}
