package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify channel on receiving below signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGHUP, syscall.SIGTERM)

	//
	go func(){
		sig := <- sigs
		fmt.Println("Received the signal", sig)
		done <- true
	}()

	go func ()  {
		for {
			select {
			case <- done:
				fmt.Println("Stopping the process due to signal")
				os.Exit(100)
			default:
				fmt.Println("waiting...")
				time.Sleep(time.Second)
			}
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}

/*
//CUSTOMISED SPECIFIC SIGNALS - SIGINT, SIGTERM, SIGHUP
func main() {
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel on receiving below signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGHUP, syscall.SIGTERM)

	go func() {
		// sig := <-sigs
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:	//here we have customized the working of SIGINT signal i.e. to log & continue
				fmt.Println("Received SIGINT (Interrupt)")	//this functionality is over-ridden. By ctrl+c won't terminate
				continue								//this process.
			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM (Terminate)")
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP (Hangup)")
			case syscall.SIGUSR1:
				fmt.Println("Received SIGNUSR1 (User defined Signal 1)")
				fmt.Println("User define function is executed")
				continue
			}
			fmt.Println("Graceful exit.")
			os.Exit(0)
		}
	}()

	for {
		fmt.Println("waiting...")
		time.Sleep(time.Second * 5)
	}
}*/

/*
//GRACEFUL EXIT ON RECEIVING SPECIFIC SIGNALS - USED FOR GRACEFUL SHUTDOWN
// & BACKING UP DATA
func main() {
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel on receiving below signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGHUP, syscall.SIGTERM)

	go func() {
		sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Received SIGINT (Interrupt)")
			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM (Terminate)")
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP (Hangup)")
			case syscall.SIGUSR1:
				fmt.Println("Received SIGNUSR1 (User defined Signal 1)")
				fmt.Println("User define function is executed")
			}
			fmt.Println("Graceful exit.")
			os.Exit(0)
	}()

	for {
		fmt.Println("waiting...")
		time.Sleep(time.Second * 5)
	}
}
*/

/* COPIED ASHISH CODE
func main() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGHUP)

	// go func() {
	// 	sig := <-sigs
	// 	fmt.Println("Received signal:", sig)
	// 	done <- true
	// }()

	go func() {

		// for {
		// 	select {
		// 	case <-done:
		// 		fmt.Println("Stopping work due to signal.")
		// 		// os.Exit(0)
		// 		return
		// 	default:
		// 		fmt.Println("Working...")
		// 		time.Sleep(time.Second)
		// 	}
		// }
		sig := <-sigs
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Received SIGINT (Interrupt)")
			// case syscall.SIGTERM:
			// 	fmt.Println("Received SIGTERM (Terminate)")
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP (Hangup)")
			case syscall.SIGUSR1:
				fmt.Println("Received SIGNUSR1 (User defined Signal 1)")
				fmt.Println("User define function is executed")
				// continue
			}
			// fmt.Println("Graceful exit.")
			// os.Exit(0)
		}
	}()

	// Simulate some work
	// fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}

// tasklist - List of all processes on Windows
// taskkill /F /PID <PID>  : taskkill /F /PID 12345
// Stop-Process -Id 12345 -Force
*/
