package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"aeropath/internal/auth"
)

// 📖 DDIA Chapitre 11 : "Stream Processing"
//    L'authentification est un "cross-cutting concern" :
//    elle est vérifiée à chaque requête, avant la logique métier.
//    C'est implémenté comme un middleware Gin (AuthMiddleware).
//    Si l'auth échoue, la requête est rejetée avant d'atteindre
//    le handler. C'est le pattern "Circuit Breaker" du Chapitre 1.

type registerRequest struct {
	Email    string `json:"email"    binding:"required,email" example:"pilote@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"monmotdepasse"`
	Lang     string `json:"lang"     binding:"omitempty,oneof=fr en" example:"fr"`
}

type loginRequest struct {
	Email    string `json:"email"    binding:"required,email" example:"pilote@example.com"`
	Password string `json:"password" binding:"required" example:"monmotdepasse"`
}

// RegisterHandler
//
//	@Summary		Inscription
//	@Description	Crée un nouveau compte étudiant et retourne un token JWT
//	@Tags			Authentification
//	@Accept			json
//	@Produce		json
//	@Param			body	body		registerRequest	true	"Email, mot de passe et langue"
//	@Success		201		{object}	map[string]string	"token JWT"
//	@Failure		400		{object}	map[string]string	"erreur de validation"
//	@Failure		409		{object}	map[string]string	"email déjà utilisé"
//	@Router			/auth/register [post]
func RegisterHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req registerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Lang == "" {
			req.Lang = "fr"
		}
		token, err := svc.Register(req.Email, req.Password, req.Lang)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("token", token, 86400, "/", "", false, true)
		c.JSON(http.StatusCreated, gin.H{"token": token})
	}
}

// LoginHandler
//
//	@Summary		Connexion
//	@Description	Authentifie un étudiant et retourne un token JWT
//	@Tags			Authentification
//	@Accept			json
//	@Produce		json
//	@Param			body	body		loginRequest	true	"Email et mot de passe"
//	@Success		200		{object}	map[string]string	"token JWT"
//	@Failure		400		{object}	map[string]string	"erreur de validation"
//	@Failure		401		{object}	map[string]string	"email ou mot de passe incorrect"
//	@Router			/auth/login [post]
func LoginHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := svc.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("token", token, 86400, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

type updateLangRequest struct {
	Lang string `json:"lang" binding:"required,oneof=fr en" example:"en"`
}

// UpdateLangHandler
//
//	@Summary		Changer la langue
//	@Description	Modifie la langue préférée de l'étudiant connecté
//	@Tags			Profil
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		updateLangRequest	true	"Nouvelle langue (fr ou en)"
//	@Success		200		{object}	map[string]string	"langue mise à jour"
//	@Failure		400		{object}	map[string]string	"erreur de validation"
//	@Router			/api/me/lang [patch]
// LogoutHandler
//
//	@Summary		Déconnexion
//	@Description	Supprime le cookie JWT httpOnly (pas de logique côté serveur)
//	@Tags			Authentification
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	map[string]string	"déconnecté"
//	@Router			/auth/logout [post]
func LogoutHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie(
			"token",
			"",
			-1,
			"/",
			"",
			false, // secure=false en dev (true en prod quand HTTPS)
			true,  // httpOnly
		)
		c.JSON(http.StatusOK, gin.H{"message": "déconnecté"})
	}
}

func UpdateLangHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req updateLangRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		studentID := c.GetString("student_id")
		if err := svc.UpdateLang(studentID, req.Lang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"lang": req.Lang})
	}
}
