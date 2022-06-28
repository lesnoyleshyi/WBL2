package shell

import (
	"bufio"
	"bytes"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// Command struct
type Command struct {
	cmd  string
	args []string
}

// Shell struct
type Shell struct {
	commands []Command
	input    *bytes.Buffer
	output   *bytes.Buffer
	out      io.Writer
}

// New return *Shell instance
func New() *Shell {
	return &Shell{}
}

// GetCommands return list of commands
func (s *Shell) GetCommands() {
	s.commands = make([]Command, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("shell$ ") //nolint:forbidigo
	line, _, _ := reader.ReadLine()

	for _, cmd := range strings.Split(string(line), "|") {
		cmd = strings.Trim(cmd, " ")
		command := strings.Split(string(cmd), " ")
		if command[0] == "top" {
			command = append(command, "-b")
		}
		s.commands = append(s.commands, Command{cmd: command[0], args: command[1:]})
	}
}

// ChangeDir changes directory
func (s *Shell) ChangeDir(cmd Command) error {
	path := cmd.args[0]
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("error when change dir: %v", err)
	}

	return nil
}

// PWD shows current directory
func (s *Shell) PWD() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error when get PWD: %v", err)
	}
	fmt.Fprintln(s.out, dir)

	return nil
}

// Echo print info to stdin
func (s *Shell) Echo(cmd Command) {
	fmt.Fprintln(s.out, strings.Join(cmd.args, " "))
}

// Kill process
func (s *Shell) Kill(cmd Command) error {
	pid, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return fmt.Errorf("incorrect PID: %v", err)
	}
	err = syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		return fmt.Errorf("can't kill process %d: %v", pid, err)
	}

	return nil
}

// PS shows list of running processes
func (s *Shell) PS(cmd Command) error {
	processes, err := ps.Processes()
	if err != nil {
		return fmt.Errorf("can't get processes list: %v", err)
	}

	for _, proc := range processes {
		fmt.Fprintf(s.out, "%d\t%s\n", proc.Pid(), proc.Executable())
	}

	return nil
}

// Exec executes binary files
func (s *Shell) Exec(cmd Command) error {
	execCmd := exec.Command(cmd.cmd, cmd.args...)

	execCmd.Stdout = s.out
	execCmd.Stdin = s.input

	if err := execCmd.Start(); err != nil {
		return fmt.Errorf("can't start executable: %v", err)
	}

	sigquit := make(chan os.Signal, 1)
	go func() {
		signal.Notify(sigquit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigquit:
			execCmd.Process.Kill()
		}
	}()

	if err := execCmd.Wait(); err != nil {
		return fmt.Errorf("executable failed: %v", err)
	}

	close(sigquit)
	defer signal.Reset()

	return nil
}

//ExecCommands select mode fo command to run
func (s *Shell) ExecCommands() error {
	var err error

	s.input = new(bytes.Buffer)
	s.output = new(bytes.Buffer)
	s.out = s.output
	for i, cmd := range s.commands {

		if i == len(s.commands)-1 {
			s.out = os.Stdin
		}

		switch cmd.cmd {
		case "cd":
			err = s.ChangeDir(cmd)
		case "pwd":
			err = s.PWD()
		case "echo":
			s.Echo(cmd)
		case "kill":
			err = s.Kill(cmd)
		case "ps":
			err = s.PS(cmd)
		case "exit":
			os.Exit(0)
		default:
			err = s.Exec(cmd)
		}

		if err != nil {
			return fmt.Errorf("error when exec command: %v", err)
		}

		s.input.ReadFrom(s.output)
	}
	return nil
}

// Run launch shell
func (s *Shell) Run() {
	for {
		s.GetCommands()
		if s.commands[0].cmd != "" {
			if err := s.ExecCommands(); err != nil {
				log.Println(fmt.Errorf("error: %v", err))
			}
		}
	}
}
