package lexec_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/reconquest/lexec-go"
	"github.com/seletskiy/hierr"
)

func ExampleLoggedExec() {
	logger := log.New(os.Stdout, `LOG: `, 0)

	cmd := lexec.New(lexec.Loggerf(logger.Printf), `wc`, `-l`)

	cmd.SetStdin(bytes.NewBufferString("1\n2\n3\n"))

	err := cmd.Run()
	if err != nil {
		log.Fatalln(hierr.Errorf(
			err,
			`can't run example command`,
		))
	}

	stdout, err := ioutil.ReadAll(cmd.GetStdout())
	if err != nil {
		log.Fatalln(hierr.Errorf(
			err,
			`can't read command stdout`,
		))
	}

	fmt.Printf("OUT: %s\n", stdout)

	// Output:
	// LOG: <stdout> {wc} 3
	// OUT: 3
}
