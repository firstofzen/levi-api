package error

import (
	"errors"
	"fmt"
)

type DBError struct{}

func (d DBError) Connect(mess string) error {
	return errors.New(fmt.Sprintf("error when connect to db | mess : { %s }", mess))
}
func (d DBError) InitTransError(mess string) error {
	return errors.New(fmt.Sprintf("error when init trans | mess : { %s }", mess))
}
func (d DBError) TransError(mess string) error {
	return errors.New(fmt.Sprintf("error when trans | mess : { %s }", mess))
}
