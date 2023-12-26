package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

type shell struct {
	out io.Writer
}

func (app *shell) run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		input := scanner.Text()
		if strings.Contains(input, "&") {
			args := strings.Split(input, "&")
			forks := args[:len(args)-1]

			if len(args) <= 2 {
				if err := app.fork(args[0]); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
				if err := app.execOrPipe(args[1]); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
			} else {
				for i := 0; i < len(forks)-1; i++ {
					if err := app.fork(forks[i]); err != nil {
						fmt.Fprintf(os.Stderr, "%v\n", err)
					}
				}
				if err := app.execOrPipe(forks[len(forks)-1]); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
			}
		} else if strings.Contains(input, "|") {
			if err := app.execPipe(input); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			app.out = os.Stdout
			if err := app.execCommand(input); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		}

		path, _ := filepath.Abs(".")
		fmt.Printf("%s $: ", path)
		scanner.Scan()
	}
}

func (app *shell) execOrPipe(input string) error {
	input = strings.TrimLeft(input, " ")
	if strings.Contains(input, "|") {
		if err := app.execPipe(input); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	} else {
		app.out = os.Stdout
		if err := app.execCommand(input); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
	return nil
}

func (app *shell) fork(input string) error {
	errors := make(chan error, 1)
	go func() { errors <- app.execOrPipe(input) }()
	if err := <-errors; err != nil {
		return err
	}
	return nil
}

func (app *shell) execCommand(command string) error {
	c := strings.Split(command, " ")

	switch c[0] {
	case "cd":
		if len(c) < 2 {
			dir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			return os.Chdir(dir)
		}
		return os.Chdir(c[1])
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Fprintln(app.out, pwd)
		return nil
	case "echo":
		for i := 1; i < len(c); i++ {
			fmt.Fprint(app.out, c[i], " ")
		}
		fmt.Fprintln(app.out)
		return nil
	case "kill":
		if len(c) < 2 {
			return fmt.Errorf("kill: not enough arguments")
		}

		pid, err := strconv.Atoi(c[1])
		if err != nil {
			p, err := ps.Processes()
			if err != nil {
				return err
			}
			for _, v := range p {
				if v.Executable() == c[1] {
					pid = v.Pid()
					break
				}
			}
			if pid == 0 {
				return fmt.Errorf("kill: can't find process with PID 0")
			}
		}
		p, err := os.FindProcess(pid)
		if err != nil {
			return err
		}

		return p.Kill()
	case "ps":
		p, err := ps.Processes()
		if err != nil {
			return err
		}
		for _, v := range p {
			fmt.Fprintf(app.out, "%d\t%s\n", v.Pid(), v.Executable())
		}
		return nil
	case "exec":
		if len(c) < 2 {
			return fmt.Errorf("exec: not enough arguments")
		}
		binary, err := exec.LookPath(c[1])
		if err != nil {
			return err
		}
		env := os.Environ()
		return syscall.Exec(binary, c[1:], env)
	case "exit":
		fmt.Fprint(app.out, "Exiting from shell\n")
		os.Exit(0)
	case "":
		fmt.Fprintf(app.out, "")
	default:
		return fmt.Errorf("command not found: %s", c[0])
	}
	return nil
}

func (app *shell) execPipe(command string) error {
	c := strings.Split(command, " | ")
	if len(c) < 2 {
		return fmt.Errorf("pipe: not enough commands: '%v'", c)
	}

	var b bytes.Buffer
	for i := 0; i < len(c); i++ {
		com := exec.Command(c[i])
		commArgs := strings.Split(c[i], " ")
		if len(commArgs) > 1 {
			com = exec.Command(commArgs[0], commArgs[1:]...)
		}

		com.Stdin = bytes.NewReader(b.Bytes())
		b.Reset()
		com.Stdout = &b

		err := com.Start()
		if err != nil {
			return err
		}
		err = com.Wait()
		if err != nil {
			return err
		}
	}

	fmt.Fprint(app.out, b.String())

	return nil
}

func main() {
	var app shell
	app.run()
}
