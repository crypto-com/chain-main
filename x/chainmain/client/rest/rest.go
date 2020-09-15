package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterRoutes registers chainmain-related REST handlers to a router
func RegisterRoutes(ctx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding
}
