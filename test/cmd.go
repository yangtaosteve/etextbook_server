package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func eatOutput(scanner *bufio.Scanner) {
	fmt.Println("Eating output")
	go func() {
		for scanner.Scan() {
			fmt.Println("==>", scanner.Text())
		}
	}()
}

func main() {
	cmd := exec.Command("python3", "-m", "pdb", "python/test.py")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd_list := []string{"next", "next", "next"}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for _, c := range cmd_list {
		fmt.Println(c)
		fmt.Fprintln(stdin, c)
		eatOutput(scanner)
	}

	stdin.Close()
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}

}
