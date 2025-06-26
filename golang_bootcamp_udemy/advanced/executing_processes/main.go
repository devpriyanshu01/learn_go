package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	pr, pw := io.Pipe()

	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("foo is good\nbar\nbaz\n"))
	}()

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("output:", string(output))
}

// func main() {
// 	cmd := exec.Command("printenv", "SHELL")

// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("Error running the cmd", err)
// 		return
// 	}
// 	fmt.Println("output:", string(output))
// }

//HOW TO WAIT FOR A PROCESS OR KILL A PROCESS IF TAKING LONG TIME.
// func main() {
// 	cmd := exec.Command("sleep", "5")

// 	err := cmd.Start()
// 	if err != nil {
// 		fmt.Println("Error Starting:", err)
// 		return
// 	}

// 	// err = cmd.Wait()
// 	time.Sleep(2 * time.Second)
// 	err = cmd.Process.Kill()	//kill a process even before it processes.
// 	if err != nil {
// 		fmt.Println("Error Killing Process:", err)
// 		return
// 	}

// 	// fmt.Println("CMD EXECUTED, PROCESS COMPLETE!!!")
// 	fmt.Println("PROCESS KILLED")
// }

//RUNNING GREP COMMAND FROM GO CODE.
// func main(){
// 	cmd := exec.Command("grep", "foo")

// 	//cmd input from where it will read.
// 	cmd.Stdin = strings.NewReader("foo is the name of the dog.\n Isn't this name cute.\nYes it is. foo")

// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}
// 	fmt.Println("output of the cmd is:\n", string(output))
// }

// running echo command
// func main() {
// 	cmd := exec.Command("echo", "Hello World")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}
// 	fmt.Println("output:", string(output))
// }
