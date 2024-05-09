package http

import (
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/pkg/errors"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func ErrLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := negroni.NewResponseWriter(w)

		next.ServeHTTP(lrw, r)

		if lrw.Status() < 200 || lrw.Status() > 299 {
			log.Println(http.StatusText(lrw.Status()))
		}
		return
	})
}

func AuthParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDCookie, err := r.Cookie(helpers.UserIDHeader)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "can't get user_id cookie", http.StatusInternalServerError)
			return
		}

		tokenCookie, err := r.Cookie(helpers.UserIDHeader)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "can't get token cookie", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r.WithContext(helpers.SetTokenAndUserIDToCtx(r.Context(), userIDCookie.Value, tokenCookie.Value)))
	})
}

// TODO: исправить
//func Metrics(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		method := r.Method
//		defer func() {
//			go common.HandlerRPS.WithLabelValues(method, status.Code(err).String()).Inc()
//		}()
//
//		start := time.Now()
//		resp, err = handler(ctx, req)
//
//		go common.HandlerRT.WithLabelValues(method).Observe(time.Since(start).Seconds())
//		return resp, err
//	})
//}
