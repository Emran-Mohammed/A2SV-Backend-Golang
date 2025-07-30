package infrastructure_test


import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"task_manager/Infrastructure"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var jwtSecret = []byte("test-secret")

// Helper to generate token
func generateToken(t *testing.T, role string, expireNow bool) string {
	claims := jwt.MapClaims{
		"user_id":  "123",
		"username": "testuser",
		"role":     role,
		"exp":      time.Now().Add(5 * time.Minute).Unix(),
	}
	if expireNow {
		claims["exp"] = time.Now().Add(-1 * time.Minute).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	assert.NoError(t, err)
	return tokenString
}

func TestRequireAuthMiddleware(t *testing.T) {
	router := gin.Default()
	router.Use(infrastructure.RequireAuth(jwtSecret))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	t.Run("missing token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("invalid format", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "InvalidToken")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("expired token", func(t *testing.T) {
		token := generateToken(t, "user", true)
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("valid token", func(t *testing.T) {
		token := generateToken(t, "user", false)
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
func TestRequireRoleMiddleware(t *testing.T) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("role", "user") // simulate a non-admin
	})
	router.Use(infrastructure.RequireRole("admin"))
	router.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin Access"})
	})

	t.Run("access denied", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/admin", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusForbidden, resp.Code)
	})

	// change role to admin
	router = gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("role", "admin") // simulate an admin
	})
	router.Use(infrastructure.RequireRole("admin"))
	router.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin Access"})
	})

	t.Run("access granted", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/admin", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
