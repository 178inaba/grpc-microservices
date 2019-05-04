package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/178inaba/grpc-microservices/front/session"
	"github.com/178inaba/grpc-microservices/front/support"
	pbuser "github.com/178inaba/grpc-microservices/proto/user"
	"github.com/rs/xid"
)

const xRequestIDKey = "X-Request-Id"

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(xRequestIDKey)
		if traceID == "" {
			traceID = newTraceID()
		}

		ctx := support.AddTraceIDToContext(r.Context(), traceID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func newTraceID() string {
	return xid.New().String()
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		defer func() {
			// TODO Replace zap.
			log.Print(start, r.Method, r.URL.String(), lrw.statusCode)
		}()

		// TODO Using lrw?
		next.ServeHTTP(w, r)
	})
}

func NewAuthentication(userClient pbuser.UserServiceClient, sessionStore session.Store) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			sessionID := session.GetSessionIDFromRequest(r)
			v, ok := sessionStore.Get(sessionID)
			if !ok {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			userID, ok := v.(uint64)
			if !ok {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			ctx := r.Context()
			resp, err := userClient.FindUser(ctx, &pbuser.FindUserRequest{UserId: userID})
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			ctx = support.AddUserToContext(ctx, resp.User)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
