package tevalis

import (
	"fmt"
)

// APIError represents a non-2xx response from the Tevalis API.
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	if len(e.Body) == 0 {
		return fmt.Sprintf("tevalis api error: status %d", e.StatusCode)
	}
	return fmt.Sprintf("tevalis api error: status %d, body: %s", e.StatusCode, string(e.Body))
}
