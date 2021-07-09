package authenticator

import (
	"context"
	"net/http"

	"github.com/consensys/quorum-key-manager/src/infra/log"

	"github.com/consensys/quorum-key-manager/src/auth/manager"
	"github.com/consensys/quorum-key-manager/src/auth/types"
)

// Middleware synchronize authentication
type Middleware struct {
	authenticator Authenticator
	policyMngr    manager.Manager
	logger        log.Logger
}

func NewMiddleware(policyMngr manager.Manager, logger log.Logger, authenticators ...Authenticator) *Middleware {
	return &Middleware{
		authenticator: First(authenticators...),
		policyMngr:    policyMngr,
		logger:        logger,
	}
}

func (mid *Middleware) Then(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		mid.ServeHTTP(rw, req, h)
	})
}

func (mid *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.Handler) {
	ctx := req.Context()

	// Authenticate request
	info, err := mid.authenticator.Authenticate(req)
	if err != nil {
		OnError(rw, req, err)
		return
	}

	if info != nil {
		// If authentication succeeded then sets the system:authenticated group
		info.Groups = append(types.AuthenticatedUser.Groups, info.Groups...)
	} else {
		// If no authentication then sets info to anonymous user
		info = types.AnonymousUser
	}

	policies := mid.getUserPolicies(ctx, info)
	mid.logger.With("policies", policies).Debug("request successfully authenticated")

	ctx = WithUserContext(ctx, NewUserContext(info))

	// Serve next
	next.ServeHTTP(rw, req.WithContext(ctx))
}

func (mid *Middleware) getUserPolicies(ctx context.Context, info *types.UserInfo) []types.Policy {
	// Retrieve policies associated to user info
	var policies []types.Policy
	for _, groupName := range info.Groups {
		group, err := mid.policyMngr.Group(ctx, groupName)
		if err != nil {
			mid.logger.WithError(err).With("group", groupName).Debug("could not load group")
			continue
		}

		for _, policyName := range group.Policies {
			policy, err := mid.policyMngr.Policy(ctx, policyName)
			if err != nil {
				mid.logger.WithError(err).With("policy", groupName).Debug("could not load policy")
				continue
			}
			policies = append(policies, *policy)
		}
	}

	// Create resolver
	return policies
}

func OnError(w http.ResponseWriter, _ *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusUnauthorized)
}
