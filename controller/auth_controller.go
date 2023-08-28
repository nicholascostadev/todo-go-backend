package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nicholascostadev/todo-backend/constant"
	"github.com/nicholascostadev/todo-backend/service"
)

type NewAuthController struct{}

var authService = service.NewAuthService{}

type RegisterUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (T *NewAuthController) RegisterUser(c *fiber.Ctx) error {
	var requestBody RegisterUserRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User registration Data"})
	}

	newUser, err := authService.RegisterUser(service.RegisterUserData{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.JSON(newUser)
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionOutput struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (T *NewAuthController) LoginUser(c *fiber.Ctx) error {
	var requestBody LoginInput

	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(requestBody)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid login request data"})
	}

	session, err := authService.LoginUser(service.LoginUserData{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		fmt.Println("error: ", err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Wrong email or password"})
	}

	returnableSession := SessionOutput{
		ID:        session.ID,
		UserID:    session.UserID,
		Token:     session.Token,
		CreatedAt: session.CreatedAt,
		UpdatedAt: session.UpdatedAt,
		ExpiresAt: session.ExpiresAt,
	}

	sessionCookie := new(fiber.Cookie)
	sessionCookie.Name = "sessionId"
	sessionCookie.Value = strconv.Itoa(int(session.ID))
	sessionCookie.Expires = constant.GenerateSessionExpiresAt()
	sessionCookie.Path = "/"
	sessionCookie.HTTPOnly = true

	c.Cookie(sessionCookie)

	c.Set("Authorization", fmt.Sprintf("Bearer %v", session.ID))

	return c.JSON(returnableSession)
}

type GetSessionByIdOutput struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"userId"`
	Token  string `json:"token"`
}

func (T *NewAuthController) GetSessionById(c *fiber.Ctx) error {
	sessionId := c.Cookies("sessionId", "")

	fmt.Println(fmt.Sprintf("Session is: %s", sessionId))
	fmt.Println(sessionId)

	if sessionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid session Id"})
	}

	convertedSessionId, err := strconv.ParseUint(sessionId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid session Id type"})
	}

	session, err := authService.GetSessionById(uint(convertedSessionId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(GetSessionByIdOutput{
		ID:     session.ID,
		UserID: session.UserID,
		Token:  session.Token,
	})
}
