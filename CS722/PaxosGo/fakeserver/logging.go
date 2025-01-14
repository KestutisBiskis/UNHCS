package fakeserver

import "fmt"

const ( // AI making my life easier
	Reset      = "\033[0m"
	Bold       = "\033[1m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Magenta    = "\033[35m"
	Cyan       = "\033[36m"
	White      = "\033[37m"
	BoldRed    = "\033[1;31m"
	BoldGreen  = "\033[1;32m"
	BoldYellow = "\033[1;33m"
	BoldBlue   = "\033[1;34m"
	BoldPurple = "\033[1;35m"
	BoldCyan   = "\033[1;36m"
	BoldWhite  = "\033[1;37m"
	BlackBg    = "\033[40m"
	RedBg      = "\033[41m"
	GreenBg    = "\033[42m"
	YellowBg   = "\033[43m"
	BlueBg     = "\033[44m"
	MagentaBg  = "\033[45m"
	CyanBg     = "\033[46m"
	WhiteBg    = "\033[47m"
)

func ColorTest() {
	// Test all colors
	fmt.Printf("%sRed%s ", Red, Reset)
	fmt.Printf("%sGreen%s ", Green, Reset)
	fmt.Printf("%sYellow%s ", Yellow, Reset)
	fmt.Printf("%sBlue%s ", Blue, Reset)
	fmt.Printf("%sMagenta%s ", Magenta, Reset)
	fmt.Printf("%sCyan%s ", Cyan, Reset)
	fmt.Printf("%sWhite%s\n", White, Reset)
	fmt.Printf("%sRed%s ", BoldRed, Reset)
	fmt.Printf("%sGreen%s ", BoldGreen, Reset)
	fmt.Printf("%sYellow%s ", BoldYellow, Reset)
	fmt.Printf("%sBlue%s ", BoldBlue, Reset)
	fmt.Printf("%sMagenta%s ", BoldPurple, Reset)
	fmt.Printf("%sCyan%s ", BoldCyan, Reset)
	fmt.Printf("%sWhite%s\n", BoldWhite, Reset)
	fmt.Printf("%sRed%s ", RedBg, Reset)
	fmt.Printf("%sGreen%s ", GreenBg, Reset)
	fmt.Printf("%sYellow%s ", YellowBg, Reset)
	fmt.Printf("%sBlue%s ", BlueBg, Reset)
	fmt.Printf("%sMagenta%s ", MagentaBg, Reset)
	fmt.Printf("%sCyan%s ", CyanBg, Reset)
	fmt.Printf("%sWhite%s ", WhiteBg, Reset)
	fmt.Printf("%sBlack%s\n", BlackBg, Reset)
}

/**
 * Logger which makes things easier
 * since these are pretty functions
 * for repetitive logging tasks
 */

type Logger struct{}

func (l *Logger) Basic(message string) {
	fmt.Printf("%s[>]%s %s\n", Cyan, Reset, message)
}

func (l *Logger) Status(message string) {
	fmt.Printf("%s[@]%s %s\n", Magenta, Reset, message)
}

func (l *Logger) Error(message string) {
	fmt.Printf("%s[!]%s %s\n", BoldRed, Reset, message)
}

func (l *Logger) ClientAdd(client *FakeClient) {
	fmt.Printf("%s[+]%s Client %d added\n", Green, Reset, client.ID)
}

func (l *Logger) ClientRemove(client *FakeClient) {
	fmt.Printf("%s[-]%s Client %d removed\n", Red, Reset, client.ID)
}

func (l *Logger) DropPacket(from, to int) {
	fmt.Printf("%s[:]%s Dropped packet from %d to %d\n", Yellow, Reset, from, to)
}

func (l *Logger) Send(from, to int, latency float64) {
	fmt.Printf("%s[~]%s Sent message from %d to %d, latency: %.2f\n", Blue, Reset, from, to, latency)
}

func (l *Logger) Receive(id, size int) {
	fmt.Printf("%s[|]%s %d received message of size %d bytes\n", Blue, Reset, id, size)
}

func (l *Logger) Broadcast(from int) {
	fmt.Printf("%s[=]%s Broadcast from %d\n", Blue, Reset, from)
}

func (l *Logger) Important(message string) {
	fmt.Printf("%s[#]%s %s%s%s\n", BoldRed, Reset, Bold, message, Reset)
}
