package errorTypes

import "fmt"

type NoRowsFoundError struct {
	Msg string
}

func (e NoRowsFoundError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}
