package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
*/

// Выводит аргумент в shell
func echo(args []string) {
	str := strings.Join(args, " ")
	fmt.Println(str)
}

// Выводит рабочую директорию в shell
func pwd() error {
	path, err := os.Getwd()
	fmt.Println(path)
	return err
}

// Меняет рабочую директорию
func cd(args []string) error {
	return os.Chdir(strings.Join(args, " "))
}

// Запускает новый процесс, с соответствующими аргументами
func forkExec(arg string) {
	fmt.Println("fork")
	args := strings.Split(arg, " ")
	if len(args) < 1 {
		fmt.Println("Error: Bad arguments")
	} else {
		cmd := exec.Command(args[0], args[1:]...)
		go func() {
			_ = cmd.Run()
		}()
	}
}

func psCmd() error {
	procs, err := ps.Processes()
	if err != nil {
		return err
	}
	for _, proc := range procs {
		fmt.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
	}
	return nil
}

func kill(args []string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	prc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return prc.Kill()
}

// Обрабатывает команду, введенную в shell
func execCmd(cmd string) error {
	cmd = strings.TrimSuffix(cmd, "\n")
	args := strings.Split(cmd, " ")
	switch args[0] {
	case "echo":
		echo(args[1:])
		return nil
	case "pwd":
		return pwd()
	case "cd":
		return cd(args[1:])
	case "fork-exec":
		forkExec(strings.Replace(cmd, "fork-exec ", "", 1))
	case "ps":
		return psCmd()
	case "kill":
		return kill(args)
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	return nil
}

func readCmd(sc *bufio.Scanner) string {
	fmt.Print(">")
	sc.Scan()
	cmd := sc.Text()
	return cmd
}

// цикл обработки shell до выхода
func shell() {
	sc := bufio.NewScanner(os.Stdin)
	cmd := readCmd(sc)
	for cmd != "quit" {
		if err := execCmd(cmd); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmd = readCmd(sc)
	}
}

func main() {
	shell()
}
