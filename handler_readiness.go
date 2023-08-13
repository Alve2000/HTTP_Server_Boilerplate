
package main

import "net/http"

// this function handles requests related to the service's readiness status.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

