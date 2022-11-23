package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	Echo     = "echo"
	Cd       = "cd"
	Kill     = "kill"
	Pwd      = "pwd"
	Exit     = "quit"
	Ps       = "ps"
	ExitText = "Exit"
)

type Command interface {
	Exec(args ...string) ([]byte, error)
}
type echo struct {
}

func (e *echo) Exec(args ...string) ([]byte, error) {
	return exec.Command("echo", args...).Output()
}

type cd struct {
}

func (c *cd) Exec(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return []byte(dir), nil
}

type pwd struct {
}

func (p *pwd) Exec(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

type kill struct {
}

func (k *kill) Exec(args ...string) ([]byte, error) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = process.Kill()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return []byte("killed"), nil
}

type ps struct {
}

func (p *ps) Exec(args ...string) ([]byte, error) {
	return exec.Command("ps").Output()
}

type StructShell struct {
	command Command
	output  io.Writer
}

func (s *StructShell) SetCommand(cmd Command) {
	s.command = cmd
}

func (s *StructShell) run(args ...string) {
	b, err := s.command.Exec(args...)
	_, err = fmt.Fprintln(s.output, string(b))
	if err != nil {
		fmt.Println("[err]", err.Error())
		return
	}
}

func (s *StructShell) Commands(cmds []string) {
	for _, command := range cmds {
		args := strings.Split(command, " ")

		com := args[0]
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case Echo:
			cm := &echo{}
			s.SetCommand(cm)

		case Cd:
			cm := &cd{}
			s.SetCommand(cm)

		case Kill:
			cm := &kill{}
			s.SetCommand(cm)

		case Pwd:
			cm := &pwd{}
			s.SetCommand(cm)

		case Ps:
			cm := &ps{}
			s.SetCommand(cm)

		case Exit:
			_, err := fmt.Fprintln(s.output, ExitText)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(1)
		default:
			fmt.Println("Такой команды не будет")
			continue
		}
		s.run(args...)
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	var output = os.Stdout

	shell := &StructShell{output: output}
	for {
		fmt.Print("command: ")

		if scan.Scan() {
			line := scan.Text()
			cm := strings.Split(line, " | ")

			shell.Commands(cm)
		}
	}
}
