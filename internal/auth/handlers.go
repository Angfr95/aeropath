package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type registerRequest struct {
    Email    string `json:"email"    binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
    Lang     string `json:"lang"     binding:"omitempty,oneof=fr en"`
}

type loginRequest struct {
    Email    string `json:"email"    binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func RegisterHandler(svc *Service) gin.HandlerFunc {
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
        c.JSON(http.StatusCreated, gin.H{"token": token})
    }
}

func LoginHandler(svc *Service) gin.HandlerFunc {
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
        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}

type updateLangRequest struct {
    Lang string `json:"lang" binding:"required,oneof=fr en"`
}

func UpdateLangHandler(svc *Service) gin.HandlerFunc {
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
