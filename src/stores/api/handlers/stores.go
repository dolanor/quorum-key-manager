package handlers

import (
	"net/http"

	storesmanager "github.com/consensys/quorum-key-manager/src/stores/manager"
	"github.com/gorilla/mux"
)

type StoresHandler struct {
	stores storesmanager.Manager

	secrets *SecretsHandler
	keys    *KeysHandler
	eth1    *Eth1Handler
}

// NewStoresHandler creates a http.Handler to be served on /stores
func NewStoresHandler(s storesmanager.Manager) *StoresHandler {
	return &StoresHandler{
		stores:  s,
		secrets: NewSecretsHandler(s),
		keys:    NewKeysHandler(s),
		eth1:    NewAccountsHandler(s),
	}
}

func (h *StoresHandler) Register(router *mux.Router) {
	// Create subrouter for /stores
	storesSubrouter := router.PathPrefix("/stores").Subrouter()

	// Create subrouter for /stores/{storeName}
	storeSubrouter := storesSubrouter.PathPrefix("/{storeName}").Subrouter()
	storeSubrouter.Use(StoreSelector)

	// Register secrets handler on /stores/{storeName}/secrets
	secretsSubrouter := storeSubrouter.PathPrefix("/secrets").Subrouter()
	h.secrets.Register(secretsSubrouter)

	// Register keys handler on /stores/{storeName}/keys
	keysSubrouter := storeSubrouter.PathPrefix("/keys").Subrouter()
	h.keys.Register(keysSubrouter)

	// Register eth1 handler on /stores/{storeName}/eth1
	eth1Subrouter := storeSubrouter.PathPrefix("/eth1").Subrouter()
	h.eth1.Register(eth1Subrouter)
}

func StoreSelector(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(WithStoreName(r.Context(), mux.Vars(r)["storeName"])))
	})
}
