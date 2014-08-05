package gocontroller

import (
	"io"
	"log"
	"net/http"
)

type Button int

const (
	NONE Button = iota
	UP
	DOWN
	LEFT
	RIGHT
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
			go func() {
				ch <- input{
					Button: UP,
					UserIP: req.RemoteAddr,
				}
			}()
		})
		http.HandleFunc("/buttonDOWN", func(w http.ResponseWriter, req *http.Request) {
			go func() {
				ch <- input{
					Button: DOWN,
					UserIP: req.RemoteAddr,
				}
			}()
		})
		http.HandleFunc("/buttonLEFT", func(w http.ResponseWriter, req *http.Request) {
			go func() {
				ch <- input{
					Button: LEFT,
					UserIP: req.RemoteAddr,
				}
			}()
		})
		http.HandleFunc("/buttonRIGHT", func(w http.ResponseWriter, req *http.Request) {
			go func() {
				ch <- input{
					Button: RIGHT,
					UserIP: req.RemoteAddr,
				}
			}()
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
