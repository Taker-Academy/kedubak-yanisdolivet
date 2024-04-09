package controllers

import (
	"context"
	"fmt"
	"time"
	"os"

	"example/kedubak-yanisdolivet/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
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

		fmt.Println("email = " + user.Email)
		fmt.Println("password = " + user.Password)
		fmt.Println("firstName = " + user.FirstName)
		fmt.Println("lastName = " + user.LastName)

		// Generate JWT token
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(KeyForToken())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Failed to generate token",
			})
		}

		fmt.Println("Token = " + tokenString)

		_, err = collection.InsertOne(ctx, user)
		if err != nil {
		    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		        "ok":    false,
		        "error": "Failed to insert user into database",
		    })
		}

		// Send Response
		return c.Status(fiber.StatusCreated).JSON(fiber.Map {
			"ok":	true,
			"data": fiber.Map {
				"token": tokenString,
				"user": fiber.Map {
					"email": user.Email,
					"firstName": user.FirstName,
					"lastName": user.LastName,
				},
			},
		})
	})
}


