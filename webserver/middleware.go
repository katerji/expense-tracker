package webserver

import (
	"github.com/katerji/expense-tracker/service/user"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if isAnonymousRequest(r) {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		authTokenFull := r.Header.Get("Authorization")
		const bearerPrefix = "Bearer "
		if len(authTokenFull) <= len(bearerPrefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		authToken := authTokenFull[len(bearerPrefix)+1:]
		u, err := user.GetServiceInstance().VerifyToken(authToken)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx = u.Ctx(ctx)

	})
}

func isAnonymousRequest(r *http.Request) bool {
	path := r.URL.Path
	for _, anonRoute := range anonRoutes() {
		if strings.Contains(path, anonRoute) {
			return true
		}
	}

	return false
}

func anonRoutes() []string {
	return []string{
		"/auth/register",
		"/auth/login",
		"/transaction",
	}
}
