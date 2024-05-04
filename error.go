package payhere

import "fmt"

func (e *RequestError) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}
