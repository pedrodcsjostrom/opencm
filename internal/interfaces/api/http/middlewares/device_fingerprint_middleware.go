package middlewares

import (
	"context"
	"net/http"
)

func AddDeviceFingerprint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fingerprint := r.RemoteAddr + r.UserAgent()
		ctx = context.WithValue(ctx, DeviceFingerprintKey, fingerprint)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}