package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("python", "../../../python/detect.py").Output()
	if err != nil {
		fmt.Println("Command Exec Error.")
	}

	// 実行したコマンドの結果を出力
	fmt.Printf("\n%s", string(out))
}
