package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/aliml92/realworld-gin-sqlc/db/sqlc"
)

type userResponse struct {
	User struct {
		Name string  `json:"name"`
		PhoneNumber    string  `json:"phone_number"`
		Otp      string `json:"otp"`
		OtpExpirationTime    time.Time  `json:"otp_expiration_time"`
		Token    string  `json:"token"`
	} `json:"user"`
}

func newUserResponse(user *db.User) *userResponse {
	resp := new(userResponse)
	resp.User.Name = user.Name
	resp.User.PhoneNumber = user.PhoneNumber
	resp.User.Otp = user.Otp
	resp.User.OtpExpirationTime = user.OtpExpirationTime
	token, _ := GenerateJWT(user.ID)
	resp.User.Token = token
	return resp
}

type userRegisterReq struct {
    User struct {
        Name         string `json:"name" binding:"required"`
        PhoneNumber string `json:"phoneNumber"`
    } `json:"user"`
}

func (r *userRegisterReq) bind(c *gin.Context, p *db.CreateUserParams) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	p.ID = generateID()
	p.Name = r.User.Name
	p.PhoneNumber = r.User.PhoneNumber
	return nil
}

func (s *Server) RegisterUser(c *gin.Context) {
	var (
		req userRegisterReq
		p   db.CreateUserParams
	)
	if err := req.bind(c, &p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, NewValidationError(err))
		return
	}
	user, err := s.store.CreateUser(c, p)
	if err != nil {
		if apiErr := convertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, NewValidationError(apiErr))
			return
		}
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, newUserResponse(user))
}

type generateOtp struct {
    Generate struct {
        PhoneNumber         string `json:"phoneNumber"`
       
    } `json:"generate"`
}

func (s *Server) GenerateOtp(c *gin.Context) {
	var (
		req generateOtp
		p   db.GenerateOtpParams
	)
	if err := req.bindGenerateOtp(c, &p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, NewValidationError(err))
		return
	}
	user, err := s.store.GenerateOtp(c, p)
	if err != nil {
		if apiErr := convertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, NewValidationError(apiErr))
			return
		}
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, newUserResponse(user))
}


type validateOtp struct {
    Validate struct {
        Otp         string `json:"otp"`
        PhoneNumber string `json:"phoneNumber"`
    } `json:"validate"`
}

func (r *validateOtp) bindOtp(c *gin.Context, p *db.ValidateOtpParams) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	p.PhoneNumber = r.Validate.PhoneNumber
	p.Otp = r.Validate.Otp
	return nil
}


func (r *generateOtp) bindGenerateOtp(c *gin.Context, p *db.GenerateOtpParams) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	p.PhoneNumber = r.Generate.PhoneNumber
	p.Otp = generateRandomString(4)
	p.OtpExpirationTime = time.Now().Add(time.Minute)
	return nil
}


func (s *Server) ValidateOtp(c *gin.Context) {
	var (
		req validateOtp
		p   db.ValidateOtpParams
	)
	if err := req.bindOtp(c, &p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, NewValidationError(err))
		return
	}

	user, err := s.store.ValidateOtp(c, p)


	if err != nil {
		if apiErr := convertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, NewError(apiErr))
			return
		}
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	c.JSON(http.StatusCreated, newUserResponse(user))
}