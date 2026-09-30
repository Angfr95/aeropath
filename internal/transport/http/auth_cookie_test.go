package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"aeropath/internal/auth"
	"aeropath/internal/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type cookieStudents struct {
	domain.StudentRepository
	hash string
}

func (r cookieStudents) FindByEmail(string) (*domain.Student, error) {
	return &domain.Student{ID: "test-student", PasswordHash: r.hash}, nil
}

func (r cookieStudents) Create(student *domain.Student) error {
	student.ID = "test-student"
	return nil
}

func TestSessionCookiesProtectProductionAndAllowLocalDevelopment(t *testing.T) {
	previous := gin.Mode()
	t.Cleanup(func() { gin.SetMode(previous) })
	hash, err := bcrypt.GenerateFromPassword([]byte("test-password"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	service := auth.NewService(cookieStudents{hash: string(hash)}, "test-secret")
	for _, mode := range []string{gin.ReleaseMode, gin.TestMode} {
		gin.SetMode(mode)
		for _, endpoint := range []string{"login", "register", "logout"} {
			t.Run(mode+"/"+endpoint, func(t *testing.T) {
				router := gin.New()
				router.POST("/login", LoginHandler(service))
				router.POST("/register", RegisterHandler(service))
				router.POST("/logout", LogoutHandler(service))
				req := httptest.NewRequest(http.MethodPost, "/"+endpoint,
					strings.NewReader(`{"email":"test@example.com","password":"test-password"}`))
				req.Header.Set("Content-Type", "application/json")
				response := httptest.NewRecorder()
				router.ServeHTTP(response, req)
				if response.Code != http.StatusOK && response.Code != http.StatusCreated {
					t.Fatalf("unexpected status %d", response.Code)
				}
				cookies := response.Result().Cookies()
				if len(cookies) != 1 || !cookies[0].HttpOnly || cookies[0].SameSite != http.SameSiteLaxMode || cookies[0].Secure != (mode == gin.ReleaseMode) {
					t.Fatal("session cookie security attributes are incorrect")
				}
			})
		}
	}
}
