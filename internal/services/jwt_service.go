package services

type JWTService interface {
	GenerateToken(userID, email string) (string, error)
	ValidateToken(token string) (*TokenClaims, error)
}

type TokenClaims struct {
	UserID string
	Email  string
}