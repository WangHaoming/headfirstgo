package utila

import (
	"errors"
	"fmt"
)

func PackageTest() error {
	fmt.Println("OK!")
	return errors.New("Package error")
}
