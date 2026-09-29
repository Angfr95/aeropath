package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type roleStub struct {
	admin bool
	err   error
}

func (r roleStub) IsAdmin(string) (bool, error) { return r.admin, r.err }

func TestAuthorization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	for _, tc := range []struct {
		name, user, target string
		owner, admin       bool
		err                error
		status             int
	}{
		{name: "anonymous", target: "other", status: 401},
		{name: "student cannot administer", user: "student", target: "other", status: 403},
		{name: "student cannot read another profile", user: "student", target: "other", owner: true, status: 403},
		{name: "owner can access profile", user: "student", target: "student", owner: true, status: 204},
		{name: "admin can administer", user: "admin", target: "other", admin: true, status: 204},
		{name: "admin can access another profile", user: "admin", target: "other", owner: true, admin: true, status: 204},
		{name: "lookup failure denies access", user: "admin", target: "other", admin: true, err: errors.New("database unavailable"), status: 403},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.New()
			r.Use(func(c *gin.Context) { c.Set("student_id", tc.user) })
			r.GET("/:id", authorize(roleStub{tc.admin, tc.err}, tc.owner), func(c *gin.Context) { c.Status(204) })
			w := httptest.NewRecorder()
			r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/"+tc.target, nil))
			if w.Code != tc.status {
				t.Fatalf("got %d, want %d", w.Code, tc.status)
			}
		})
	}
}
