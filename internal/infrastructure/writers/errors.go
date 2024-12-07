package writers

import "fmt"

type ImageCreatingError struct {
	msg string
}

func (e ImageCreatingError) Error() string {
	return fmt.Sprintf("Error: opening outgoing Image: %s", e.msg)
}
