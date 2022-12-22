package utils

import (
	"fmt"
	"strings"
)

func ErrorChecker(builder strings.Builder) (err error) {
	if builder.Len() > 0 {
		s := builder.String()
		s = s[:builder.Len()-2]
		err = fmt.Errorf(s)
	}
	return err
}
