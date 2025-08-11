package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TexttoMorseAndRevers(str string) string {
	count := 0
	for _, ch := range str {
		if string(ch) == " " || string(ch) == "." || string(ch) == "-" {
			count++
		}
	}
	if len(str) == count {
		return morse.ToText(str)
	} else {
		return morse.ToMorse(str)
	}

}
