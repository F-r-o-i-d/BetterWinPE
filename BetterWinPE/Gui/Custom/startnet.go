package Custom

import (
	"log"
	"os"
)

func SetStartnet(s string, workingDirectories string) error {
	err := os.Rename(s, workingDirectories+`\mount\Windows\System32\startnet.cmd`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
