package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	var com string
	fmt.Println("Enter exact, current or exit")
	fmt.Scan(&com)
	for com != "exit" {
		var format string
		switch com {
		case "exact":
			format = "15:04:05.000000000"
		case "current":
			format = "15:04:05"
		default:
			fmt.Println("Enter exact, current or exit")
			fmt.Scan(&com)
			continue
		}
		time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error has occured: %v", err)
			os.Exit(-1)
		}
		_, _ = fmt.Fprintf(os.Stdout, "%s time: %v\n", com, time.Format(format))
		fmt.Scan(&com)
	}
}
