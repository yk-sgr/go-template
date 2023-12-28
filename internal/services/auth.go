package services

import (
	"context"
	"time"

	"github.com/cohesivestack/valgo"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/yk-sgr/nexthook-api/internal/db"
	"github.com/yk-sgr/nexthook-api/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

// AuthService implementation for domain.AuthService.
type AuthService struct {
	db        *db.Queries
	jwtSecret string
}

func NewAuthService(db *db.Queries, jwtSecret string) domain.AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

// SignUp signs up a user.
func (s *AuthService) SignUp(ctx context.Context, dto *domain.SignUpRequest) (*domain.SignUpResponse, error) {
	// Validate input
	val :=
		valgo.Is(valgo.String(dto.Email, "email").Not().Empty().OfLengthBetween(domain.UserMinEmailLength, domain.UserMaxEmailLength)).
			Is(valgo.String(dto.Password, "password").Not().Empty().OfLengthBetween(domain.UserMinPasswordLength, domain.UserMaxPasswordLength)).
			Is(valgo.String(dto.Name, "name").Not().Empty().OfLengthBetween(domain.UserMinNameLength, domain.UserMaxNameLength))
	if !val.Valid() {
		return nil, domain.NewValidationError()
	}

	// Check if email is already registered
	_, err := s.db.GetUserByEmail(ctx, dto.Email)
	if err == nil {
		return nil, domain.NewAlreadyExistsError("signup/email-exists")
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	_, err = s.db.CreateUser(ctx, db.CreateUserParams{
		Email:    dto.Email,
		Password: string(password),
		Name:     dto.Name,
	})
	if err != nil {
		return nil, err
	}

	return &domain.SignUpResponse{}, nil
}

// SignIn signs in a user.
func (s *AuthService) SignIn(ctx context.Context, dto *domain.SignInRequest) (*domain.SignInResponse, error) {
	// Validate input
	val := valgo.Is(valgo.String(dto.Email, "email").Not().Empty().OfLengthBetween(3, 320)).
		Is(valgo.String(dto.Password, "password").Not().Empty().OfLengthBetween(8, 128))
	if !val.Valid() {
		return nil, domain.NewValidationError()
	}

	// Get user by email
	u, err := s.db.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.NewNotFoundError()
		}
		return nil, err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(dto.Password))
	if err != nil {
		return nil, domain.NewUnauthorizedError()
	}

	expiresAt := time.Now().Add(time.Hour * 24)

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        u.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &domain.SignInResponse{
		User:        *u.ToUser(),
		AccessToken: signedToken,
		ExpiresAt:   expiresAt,
	}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*domain.User, error) {
	t, err := jwt.ParseWithClaims(token, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(*domain.TokenClaims)
	if !ok {
		return nil, domain.NewUnauthorizedError()
	}

	// Fetch user
	if claims.ID == "" {
		return nil, domain.NewUnauthorizedError()
	}
	u, err := s.db.GetUserByID(ctx, uuid.MustParse(claims.ID))
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.NewUnauthorizedError()
		}
		return nil, err
	}

	return u.ToUser(), nil
}
