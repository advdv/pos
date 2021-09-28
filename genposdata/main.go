package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	f, err := os.Create("test_proofs.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 256; i++ {
		var buf bytes.Buffer
		sum := sha256.Sum256([]byte{byte(i)})

		chall := hex.EncodeToString(sum[:])
		cmd := exec.Command("./ProofOfSpace", "-f", "plot.dat", "prove", "0x"+chall)
		cmd.Stdout = &buf

		cmd.Run()

		scan := bufio.NewScanner(&buf)
		for scan.Scan() {
			txt := scan.Text()
			if !strings.HasPrefix(txt, "Proof: 0x") {
				continue
			}

			fmt.Fprintf(f, "%s:%s\n", chall, strings.TrimPrefix(txt, "Proof: 0x"))
		}
	}
}
