package main

import "syscall"
import "os"
import "os/exec"

func main() {
	
	binary, lookErr := exec.LookPath("ls") //we need absolute path
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"} //arguments in slice form
    //first argument is the program name

    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }

}