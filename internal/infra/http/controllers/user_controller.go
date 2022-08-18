package controllers

import (
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

const BEARER = "Bearer "

type UserController struct {
	userService         *app.UserService
	tokenService        *app.TokenService
	userValidator       *validators.UserValidator
	userLogInValidator  *validators.UserLogInValidator
	userUpdateValidator *validators.UserUpdateValidator
}

func NewUserController(u *app.UserService, rt *app.TokenService) *UserController {
	return &UserController{
		userService:         u,
		tokenService:        rt,
		userValidator:       validators.NewUserValidator(),
		userLogInValidator:  validators.NewUserLogInValidator(),
		userUpdateValidator: validators.NewUserUpdateValidator(),
	}
}

func (c *UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := c.userValidator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		savedUser, err := (*c.userService).Create(user)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		created(w, resources.MapDomainToUserDto(savedUser))

	}
}

func (c *UserController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print("eventsController FindOne ParseInt", err)
			badRequest(w, err)
			return
		}
		log.Print(id)
		authHeader := r.Header.Get("Authorization")
		token := authHeader[len(BEARER):]

		params, err := parseUrlQuery(r)
		if err != nil {
			log.Println(err)
			badRequest(w, err)
			return
		}
		user, err := (*c.tokenService).VerifyToken(token)
		if err != nil {
			log.Print(err)
			return
		}

		usr, err := (*c.userService).FindOne(user.UserId, params)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDto(usr))

	}
}
func (c *UserController) FindOneById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseUrlQuery(r)
		if err != nil {
			log.Println(err)
			badRequest(w, err)
			return
		}
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print("userController FindOneById ParseInt", err)
			badRequest(w, err)
			return
		}
		authHeader := r.Header.Get("Authorization")
		token := authHeader[len(BEARER):]

		_, err = (*c.tokenService).VerifyToken(token)
		if err != nil {
			log.Print(err)
			return
		}

		usr, err := (*c.userService).FindOne(id, params)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDto(usr))

	}
}
func (c *UserController) PaginateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseUrlQuery(r)
		if err != nil {
			log.Println(err)
			badRequest(w, err)
			return
		}

		users, err := (*c.userService).FindAll(params)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDtoCollection(users))

	}
}

func (c *UserController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		user, err := c.userUpdateValidator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}
		user.Id = id
		updatedUser, err := (*c.userService).Update(user)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}
		success(w, resources.MapDomainToUserDto(updatedUser))

	}
}

func (c *UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		err = (*c.userService).Delete(id)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		ok(w)
	}
}

func (c *UserController) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := c.userLogInValidator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		userStored, err := (*c.userService).LogIn(user)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		accessToken, err := (*c.tokenService).CreateToken(userStored)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		w.Header().Set("Authorization", accessToken)

		success(w, resources.MapDomainToTokenDto(accessToken))

	}
}

func (c *UserController) CheckAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok(w)
	}
}
