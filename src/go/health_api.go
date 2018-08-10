/*
 * This file will not be overwritten by successive calls to swagger-codegen
 * as it is skipped in .swagger-codegen-ignore.
 */

package swagger

import (
	"fmt"
	"net/http"
	"os"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, "OK - pid %d\n", os.Getpid())

	// Check connection to database(s)

	w.WriteHeader(http.StatusOK)
}
