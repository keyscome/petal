// color/color.go
package color

import "fmt"

const (
	AnsiReset   = "\033[0m"
	AnsiHeader  = "\033[0;36m"
	AnsiRunning = "\033[1;33m"
	AnsiSuccess = "\033[1;92m"
	AnsiError   = "\033[1;91m"
)

func PrintHeader(format string, a ...interface{}) {
	fmt.Printf(AnsiHeader+format+AnsiReset+"\n", a...)
}

func PrintRunning(format string, a ...interface{}) {
	fmt.Printf(AnsiRunning+format+AnsiReset+"\n", a...)
}

func PrintSuccess(format string, a ...interface{}) {
	fmt.Printf(AnsiSuccess+format+AnsiReset+"\n", a...)
}

func PrintError(format string, a ...interface{}) {
	fmt.Printf(AnsiError+format+AnsiReset+"\n", a...)
}
