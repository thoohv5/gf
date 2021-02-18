package util

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func WrapRecover(f func() error) error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %w, stack: %v", errors.New(fmt.Sprintf("%v", r)), string(debug.Stack()))
				return
			}
		}()
		err = f()
		return err
	}()
}
