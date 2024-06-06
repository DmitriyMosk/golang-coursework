package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var timeout time.Duration
		if r.URL.Path == "/api/v1/connector" || r.URL.Path == "/api/v1/graph" {
			timeout = time.Duration(viper.GetInt("analyticsTimeout")) * time.Second
		} else {
			timeout = time.Duration(viper.GetInt("resourseTimeout")) * time.Second
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

		if ctx.Err() == context.DeadlineExceeded {
			w.WriteHeader(http.StatusRequestTimeout)
			w.Write([]byte("Request Timeout"))
			return
		}
	})
}
