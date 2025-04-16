package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func())
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux
        cmd.Stdout = os.Stdout
        cmd.Run()
    }

    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok {
        value()
    } else {
        panic("Your platform is unsupported!")
    }
}

func PrintHeader() {
 	fmt.Println("       ______           ______             _    _                  _____ ")
	fmt.Println("       |  _  \\          | ___ \\           | |  (_)               _|____ |")
	fmt.Println("       | | | |___  _ __ | |_/ / __ _ _ __ | | ___ _ __   __ _   (_)   / /")
	fmt.Println("       | | | / _ \\| '_ \\| ___ \\/ _` | '_ \\| |/ / | '_ \\ / _` |        \\ \\")
	fmt.Println("       | |/ / (_) | | | | |_/ / (_| | | | |   <| | | | | (_| |   _.___/ /")
	fmt.Println("       |___/ \\___/|_| |_\\____/ \\__,_|_| |_|_|\\_\\_|_| |_|\\__, |  (_)____/ ")
	fmt.Println("                                                         __/ |           ")
	fmt.Println("                                                        |___/            ") 
}