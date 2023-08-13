
package main

import "net/http"

// This function handles requests related to the service's error
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}

