package auth

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"

    "aeropath/internal/domain"
)

type Service struct {
    repo      domain.StudentRepository
    jwtSecret string
}

func NewService(repo domain.StudentRepository, jwtSecret string) *Service {
    return &Service{repo: repo, jwtSecret: jwtSecret}
}

func (s *Service) Register(email, password, lang string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("hash password: %w", err)
    }

    student := &domain.Student{
        Email:        email,
        PasswordHash: string(hash),
        Lang:         lang,
    }
    if err := s.repo.Create(student); err != nil {
        return "", fmt.Errorf("create student: %w", err)
    }

    return GenerateToken(student.ID, s.jwtSecret)
}

func (s *Service) Login(email, password string) (string, error) {
    student, err := s.repo.FindByEmail(email)
    if err != nil {
        // On renvoie le même message que si le password est faux —
        // évite l'énumération d'emails (security best practice)
        return "", fmt.Errorf("email ou mot de passe incorrect")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(student.PasswordHash), []byte(password)); err != nil {
        return "", fmt.Errorf("email ou mot de passe incorrect")
    }

    return GenerateToken(student.ID, s.jwtSecret)
}

func (s *Service) UpdateLang(studentID, lang string) error {
    if lang != "fr" && lang != "en" {
        return fmt.Errorf("langue non supportée (fr ou en)")
    }
    return s.repo.UpdateLang(studentID, lang)
}
