package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"example/kedubak-yanisdolivet/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func KeyForToken() []byte {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	jwtKey := []byte(os.Getenv("TOKEN_SECRET"))
	return jwtKey
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userID string) string {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	secretKey := []byte(os.Getenv("SECRET"))
	FirstToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	FinalToken, err := FirstToken.SignedString(secretKey)
	if err != nil {
		return ""
	}
	return FinalToken
}

func GetPostRequest(app *fiber.App, Client *mongo.Client) {

	// POST route to handle incoming requests
	app.Post("/auth/register", func(c *fiber.Ctx) error {
		// Parse JSON request body into a struct
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		collection := Client.Database("ClusterKeduback").Collection("Users")

		ctx := context.Background()

		filter := bson.M{"email": user.Email}
		var existingUser models.User
		err := collection.FindOne(ctx, filter).Decode(&existingUser)

		if err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"ok":    false,
				"error": "Email already exists",
			})
		} else if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Internal server error",
			})
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Failed to hash password",
			})
		}
		user.Password = string(hashedPassword)

		// Set LastUpVote & CreateAr
		user.CreatedAt = time.Now()
		user.LastUpVote = time.Now().Add(-1 * time.Minute)

		fmt.Println("email = " + user.Email)
		fmt.Println("password = " + user.Password)
		fmt.Println("firstName = " + user.FirstName)
		fmt.Println("lastName = " + user.LastName)

		// Generate JWT token
		USERID := user.Id.Hex()
		token := GenerateToken(USERID)

		fmt.Println("Token = " + token)

		_, err = collection.InsertOne(ctx, user)
		if err != nil {
			fmt.Println("Failed to insert into db")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Failed to inster into the DataBase",
			})
		}
		fmt.Println("Insert into the db successfully")

		// Send Response
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok": true,
			"data": fiber.Map{
				"token": token,
				"user": fiber.Map{
					"email":     user.Email,
					"firstName": user.FirstName,
					"lastName":  user.LastName,
				},
			},
		})
	})
}
