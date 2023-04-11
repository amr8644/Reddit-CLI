package main

import (
	"log"
)

var (
	Reset = "\033[0m"

	/////////////
	// Special //
	/////////////

	Bold      = "\033[1m"
	Underline = "\033[4m"

	/////////////////
	// Text colors //
	/////////////////

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	///////////////////////
	// Background colors //
	///////////////////////

	BlackBackground  = "\033[40m"
	RedBackground    = "\033[41m"
	GreenBackground  = "\033[42m"
	YellowBackground = "\033[43m"
	BlueBackground   = "\033[44m"
	PurpleBackground = "\033[45m"
	CyanBackground   = "\033[46m"
	GrayBackground   = "\033[47m"
	WhiteBackground  = "\033[107m"
)

func Error(message string) {
	log.Println(RedBackground + White + "ERROR" + Reset + ": " + message)

}
func Success(message string) {
	log.Println(GreenBackground + White + "SUCCESS" + Reset + ": " + message)
}
func Info(message string) {
	log.Println(BlueBackground + White + "INFO" + Reset + ": " + message)
}
