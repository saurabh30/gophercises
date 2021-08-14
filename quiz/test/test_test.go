package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"bou.ke/monkey"
)

func TestTakeTest(t *testing.T) {
	monkey.Patch(fmt.Scanln, func(a ...interface{}) (n int, err error) {
		return 0, nil
	})
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	got1, _ := TakeTest(dir + "input.csv")
	want1 := 13

	if got1 != want1 {
		t.Errorf("got %q, wanted %q", got1, want1)
	}

}
