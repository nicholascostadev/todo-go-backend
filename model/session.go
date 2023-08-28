package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/nicholascostadev/todo-backend/constant"
	"gorm.io/gorm"
)

type Session struct {
	ID        uint      `json:"id"        gorm:"primaryKey"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user"`
	Token     string    `json:"token"     gorm:"not null"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;default:current_timestamp"`
}

func (s *Session) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}

type JWTData struct {
	SessionID uint   `json:"sessionId"`
	UserID    uint   `json:"userId"`
	Username  string `json:"username"`
}

type SessionClaims struct {
	SessionID uint   `json:"sessionId"`
	UserID    uint   `json:"userId"`
	Username  string `json:"username"`
	jwt.StandardClaims
}

func generateJWT(jwtData JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SessionClaims{
		SessionID: jwtData.SessionID,
		UserID:    jwtData.UserID,
		Username:  jwtData.Username,
	})

	tokenString, err := token.SignedString([]byte("mysupersecret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type CreateSessionInput struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func CreateSession(user CreateSessionInput) (Session, error) {
	sessionId := uint(uuid.New().ID())
	token, err := generateJWT(JWTData{
		SessionID: sessionId,
		UserID:    user.ID,
		Username:  user.Username,
	})
	if err != nil {
		fmt.Println("real error: ", err)
		return Session{}, errors.New("Could not generate JWT session token")
	}
	session := Session{
		ID:        sessionId,
		Token:     token,
		UserID:    user.ID,
		ExpiresAt: constant.GenerateSessionExpiresAt(),
	}

	tx := db.Create(&session)

	if tx.Error != nil {
		return Session{}, errors.New("There was an error when creating your session")
	}

	return session, nil
}

func DeleteSession(session Session) (Session, error) {
	tx := db.Delete(&session)

	if tx.Error != nil {
		return Session{}, tx.Error
	}

	return session, nil
}

func GetSessionById(id uint) (Session, error) {
	session := Session{ID: id}
	tx := db.Find(&session)

	if tx.Error != nil {
		return Session{}, tx.Error
	}

	if session.ExpiresAt.Before(time.Now()) {
		return DeleteSession(session)
	}

	return session, nil
}
