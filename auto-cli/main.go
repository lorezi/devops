package main

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	kubectl = "kubectl"
	git     = "git"
	openCl  = "opencl"
)

func main() {
	err := lookPath()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Configs in our path")

}

func lookPath() error {
	_, err := exec.LookPath(kubectl)
	if err != nil {
		return fmt.Errorf("cannot find kubectl in our PATH")
	}

	_, err = exec.LookPath(git)
	if err != nil {
		return fmt.Errorf("cannot find git in our PATH")
	}

	_, err = exec.LookPath(openCl)
	if err != nil {
		return fmt.Errorf("cannot find opencl in our PATH")
	}

	return nil
}
