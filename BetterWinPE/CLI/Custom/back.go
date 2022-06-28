package Custom

import (
	"io"
	"io/ioutil"
	"log"
)

func ChangeBackground(stding io.ReadCloser, workingDirectories string, backgroundPath string) {
	go func() {
		defer stding.Close()
		src := backgroundPath

		bytesRead, err := ioutil.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(workingDirectories+"\\mount\\Windows\\System32\\winpe.jpg", bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}
	}()
}
