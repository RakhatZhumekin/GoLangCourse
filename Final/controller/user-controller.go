package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/dto"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/helper"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
	emailConfirmationService service.EmailConfirmationService
}

func NewUserController(userService service.UserService, jwtService service.JWTService, emailConfirmationService service.EmailConfirmationService) UserController {
	return &userController {
		userService,
		jwtService,
		emailConfirmationService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	userUpdateDTO.ID = id
	okText := "OK!"

	if userUpdateDTO.Name == "" {
		userUpdateDTO.Name = c.userService.Profile(fmt.Sprintf("%v", claims["user_id"])).Name
	}

	if userUpdateDTO.Email == "" {
		userUpdateDTO.Email = c.userService.Profile(fmt.Sprintf("%v", claims["user_id"])).Email
		userUpdateDTO.Verified = c.userService.Profile(fmt.Sprintf("%v", claims["user_id"])).Verified
	} else {
		rand.Seed(time.Now().UnixNano())
		Rn = rand.Intn(100000)

		err := c.emailConfirmationService.SendVerificationCode(Rn, userUpdateDTO.Email)

		if err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		okText = "OK! Check your email for verification code"
		userUpdateDTO.Verified = false
	}
	
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, okText, u)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}