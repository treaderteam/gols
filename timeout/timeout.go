package timeout

import (
	"time"
)

// WithTimeout executes given function whitin given timeout
// return false if function not executed til timeout
func WithTimeout(fn func(...interface{}), out time.Duration) bool {

}
