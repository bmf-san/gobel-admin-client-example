package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
	"github.com/bmf-san/gobel-admin-client-example/app/logger"
	"github.com/bmf-san/gobel-admin-client-example/app/response"
	"github.com/bmf-san/gobel-admin-client-example/app/session"
)

const (
	ctxSessID    = "session-id"
	ctxCSRFToken = "csrf-token"
)

// middelware represents the singular of middleware.
type middleware func(http.HandlerFunc) http.HandlerFunc

// Middlewares represents the plural of middleware.
type Middlewares struct {
	middlewares []middleware
}

// Asset represents the plural of middelware services.
type Asset struct {
	redisHandler *session.RedisHandler
	logger       *logger.Logger
	cookie       *cookie.Cookie
	response     *response.Response
}

// NewAsset creates a assets.
func NewAsset(redisHandler *session.RedisHandler, logger *logger.Logger, cookie *cookie.Cookie, response *response.Response) Asset {
	return Asset{
		redisHandler: redisHandler,
		logger:       logger,
		cookie:       cookie,
		response:     response,
	}
}

// NewMiddlewares creates a middlewares.
func NewMiddlewares(mws ...middleware) Middlewares {
	return Middlewares{
		middlewares: append([]middleware(nil), mws...),
	}
}

// Then handles chaining middlewares.
func (mws Middlewares) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range mws.middlewares {
		h = mws.middlewares[len(mws.middlewares)-1-i](h)
	}

	return h
}

// Session is a middleware for session.
func (a *Asset) Session(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessID, err := a.cookie.GetSession(w, r)
		if err != nil {
			if err == http.ErrNoCookie {
				sessID, err = a.redisHandler.SetSession()
				if err != nil {
					a.logger.Error(fmt.Sprintf("An unexpected condition has occured:%v", err))
					http.Redirect(w, r, "/signin", http.StatusSeeOther)
				}

				a.cookie.SetSession(w, r, []byte(sessID))
			}
		}

		sessID, err = a.redisHandler.GetSession(sessID)
		if err != nil {
			sessID, err = a.redisHandler.SetSession()
			if err != nil {
				a.logger.Error(fmt.Sprintf("An unexpected condition has occured:%v", err))
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
			}

			a.cookie.SetSession(w, r, []byte(sessID))
		}

		ctx := context.WithValue(r.Context(), ctxSessID, sessID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Auth is a middleware for authentication.
func (a *Asset) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginSessID, err := a.cookie.GetLoginSession(w, r)
		if err != nil {
			a.logger.Error(fmt.Sprintf("An unexpected condition has occured:%v", err))
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}

		_, err = a.redisHandler.GetLoginSession(loginSessID)
		if err != nil {
			a.logger.Error(fmt.Sprintf("An unexpected condition has occured:%v", err))
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}

		next.ServeHTTP(w, r)
	}
}

// CSRF is a middelware for CSRF.
func (a *Asset) CSRF(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			token, err := a.redisHandler.SetCSRFToken()
			if err != nil {
				a.logger.Error(err.Error())
				a.response.Error(w, http.StatusInternalServerError)
				return
			}
			ctx := context.WithValue(r.Context(), ctxCSRFToken, token)
			next.ServeHTTP(w, r.WithContext(ctx))
			return

		case http.MethodPost:
			r.ParseForm()
			token := r.Form.Get("csrf_token")
			if token == "" {
				a.logger.Error("CSRF token is invalid")
				a.response.Error(w, http.StatusInternalServerError)
				return
			}

			if err := a.redisHandler.HasCSRFToken(token); err != nil {
				a.logger.Error(err.Error())
				a.response.Error(w, http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, r)
			return
		default:
			next.ServeHTTP(w, r)
			return
		}
	}
}
