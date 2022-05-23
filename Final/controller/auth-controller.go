package controller

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/dto"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/entity"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/helper"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/service"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

var Rn = 0

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RegistrationConfirm(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
	userService service.UserService
	emailConfirmationService service.EmailConfirmationService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService, userService service.UserService, emailConfirmationService service.EmailConfirmationService) AuthController {
	return &authController{
		authService,
		jwtService,
		userService,
		emailConfirmationService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		if (v.Verified) {
			generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			v.Token = generatedToken
			response := helper.BuildResponse(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
			return
		} else {
			response := helper.BuildErrorResponse("Failed to process request", "This account has not been verified", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	response := helper.BuildErrorResponse("Please check your credentials again!", "Invalid Credentials", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token

		rand.Seed(time.Now().UnixNano())
		Rn = rand.Intn(100000)

		err := c.emailConfirmationService.SendVerificationCode(Rn, createdUser.Email)

		if err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := helper.BuildResponse(true, "OK! Check your email for verification code", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *authController) RegistrationConfirm(ctx *gin.Context) {
	var emailVerificationDTO dto.EmailVerificationDTO
	errDTO := ctx.ShouldBind(&emailVerificationDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if emailVerificationDTO.Verification == strconv.Itoa(Rn) {
		verifyingUser := c.authService.FindByEmail(emailVerificationDTO.Email)
		verifyingUser.Verified = true

		verifyingUserDTO := dto.UserUpdateDTO{}

		err := smapping.FillStruct(&verifyingUserDTO, smapping.MapFields(&verifyingUser))

		if err != nil {
			log.Fatalf("Failed to map %v", err)
		}

		verifyingUserDTO.Password = ""
		c.userService.Update(verifyingUserDTO)

		response := helper.BuildResponse(true, "OK! You can now login", verifyingUserDTO)
		ctx.JSON(http.StatusCreated, response)
	} else {
		response := helper.BuildErrorResponse("Failed to process request", "Wrong verification code", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
}