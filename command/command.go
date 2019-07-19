package command

import "fmt"

type Server struct{}

func (s *Server) Register() {
	fmt.Println("register")
}

func (s *Server) Login() {
	fmt.Println("login")
}

func (s *Server) Say() {
	fmt.Println("say")
}

func (s *Server) Watch() {
	fmt.Println("watch")
}

type Command interface {
	Execute(subCmd string)
}

type UserCommand struct {
	s *Server
}

func (uc *UserCommand) Execute(subCmd string) {
	if subCmd == "register" {
		uc.s.Register()
	} else if subCmd == "login" {
		uc.s.Login()
	}
}

type InteractCommand struct {
	s *Server
}

func (ic *InteractCommand) Execute(subCmd string) {
	if subCmd == "say" {
		ic.s.Say()
	} else if subCmd == "watch" {
		ic.s.Watch()
	}
}

func CommandTest() {
	s := new(Server)

	var c Command
	c = &UserCommand{s}
	c.Execute("login")

	c = &InteractCommand{s}
	c.Execute("say")
}
