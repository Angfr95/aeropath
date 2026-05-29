package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken crée un token JWT pour un étudiant.
//
// 🛡️ SÉCURITÉ :
//    - Expiration : 24h (pas de token permanent)
//    - Signature : HMAC-SHA256 (secret partagé)
//    - Claims : student_id (identifiant unique)
//
// 🧠 POURQUOI 24h ?
//    - Trop court : l'étudiant doit se reconnecter tout le temps
//    - Trop long : si le token est volé, l'attaquant a accès longtemps
//    24h est un bon compromis pour une appli de formation.
//
//    En production, on pourrait ajouter :
//    - Refresh token (7 jours) pour renouveler sans se reconnecter
//    - JTI (JWT ID) pour révoquer un token spécifique
func GenerateToken(studentID, secret string) (string, error) {
	claims := jwt.MapClaims{
		"student_id": studentID,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
		"iat":        time.Now().Unix(), // Issued At : quand le token a été créé
		"nbf":        time.Now().Unix(), // Not Before : pas valide avant cette date
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// RequireAuth est un middleware Gin qui vérifie le token JWT.
//
// 🛡️ SÉCURITÉ :
//    1. Vérifie que le header Authorization commence par "Bearer "
//    2. Vérifie la signature HMAC (secret partagé)
//    3. Vérifie que le token n'est pas expiré (exp)
//    4. Vérifie que le token est valide (nbf)
//    5. Extrait student_id et le met dans le contexte Gin
//
// 🧠 POURQUOI "Bearer " ?
//    Le standard HTTP dit que les tokens d'authentification
//    sont envoyés dans le header Authorization avec le type "Bearer".
//    Exemple : Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
//
//    Sans "Bearer ", on ne peut pas distinguer un token JWT
//    d'un autre type d'authentification (Basic, Digest, etc.).
func RequireAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" || len(tokenStr) < 7 || tokenStr[:7] != "Bearer " {
			c.AbortWithStatusJSON(401, gin.H{"error": "token manquant"})
			return
		}
		tokenStr = tokenStr[7:]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			// Vérifier que l'algorithme de signature est bien HMAC
			// (évite les attaques par changement d'algorithme)
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signature invalide")
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "token invalide ou expiré"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "token invalide"})
			return
		}

		// Vérifier que student_id est bien une string
		studentID, ok := claims["student_id"].(string)
		if !ok || studentID == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "token invalide"})
			return
		}

		c.Set("student_id", studentID)
		c.Next()
	}
}
