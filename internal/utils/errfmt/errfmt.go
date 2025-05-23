package errfmt

import (
	"strings"

	"github.com/Bel1Vuk/jetArrays/v2/internal/utils/is"
)

// Trace returns well formatted wrapped error trace string
func Trace(err error) string {
	if is.Nil(err) {
		return ""
	}
	return "Error trace:\n" + " - " + strings.Replace(err.Error(), ": ", ":\n - ", -1)
}
