package middlewares

import (
    "github.com/Setti7/hazel/internal/models"
    "github.com/Setti7/hazel/internal/setup"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

// Extract custom claims.
var claims struct {
    Email    string   `json:"Email"`
    Verified bool     `json:"email_verified"`
    Groups   []string `json:"Groups"`
    Scope    string   `json:"scope"`
    Name     string   `json:"Name"`
}

const (
    AUTH_ERROR_HEADER_FORMATTING    = "AUTH_ERROR_HEADER_FORMATTING"
    AUTH_ERROR_INVALID_TOKEN        = "AUTH_ERROR_INVALID_TOKEN"
    AUTH_ERROR_INVALID_CLAIMS       = "AUTH_ERROR_INVALID_CLAIMS"
    AUTH_ERROR_INVALID_UNVERIFIED_EMAIL = "AUTH_ERROR_INVALID_UNVERIFIED_EMAIL"
)

func Authorize() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        splitToken := strings.Split(authHeader, "Bearer ")

        if len(splitToken) != 2 {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": AUTH_ERROR_HEADER_FORMATTING,
                "error_message": "Authorization header should be in format: Bearer <token>"})
            return
        }

        bearerToken := strings.TrimSpace(splitToken[1])

        idToken, err := setup.IdTokenVerifier.Verify(c, bearerToken)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": AUTH_ERROR_INVALID_TOKEN,
                "error_message": "Could not verify bearer token"})
            return
        }

        if err := idToken.Claims(&claims); err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": AUTH_ERROR_INVALID_CLAIMS,
                "error_message": "Could not parse claims"})
            return
        }

        if !claims.Verified {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": AUTH_ERROR_INVALID_UNVERIFIED_EMAIL,
                "error_message": "User needs to have verified email address"})
            return
        }

        // TODO verify if user has permission to access these resources? Is this what claims.Scope is?

        user := &models.User{claims.Name, claims.Email, claims.Groups, claims.Verified}
        c.Set("user", user)

        c.Next()
    }
}
