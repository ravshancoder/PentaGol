package v1

import (
	"context"
	"strconv"

	"net/http"

	_ "github.com/PentaGol/api_getway/api/docs"
	"github.com/PentaGol/api_getway/api/handlers/models"
	pu "github.com/PentaGol/api_getway/genproto/admin"
	"github.com/PentaGol/api_getway/pkg/etc"
	l "github.com/PentaGol/api_getway/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Login admin
// @Summary login admin api
// @Description this api login admin
// @Tags Admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param        email  	path string true "email"
// @Param        password   path string true "password"
// @Succes       200		{object}	models.LoginAdmin
// Failure       500        {object}  models.Error
// Failure       400        {object}  models.Error
// Failure       404        {object}  models.Error
// @Router /v1/login/{email}/{password} [get]
func (h *handlerV1) Login(c *gin.Context) {
	var (
		email    = c.Param("email")
		password = c.Param("password")
	)
	// fmt.Println("Password: ", password, "	Email: ", email)

	res, err := h.serviceManager.AdminService().GetByEmail(
		context.Background(), &pu.EmailReq{
			Email: email,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error GetByEmail login func", l.Error(err))
		return
	}

	if !etc.CheckPasswordHash(password, res.Password) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Passwrod or Email error",
		})
		return
	}

	h.jwtHandler.Iss = "user"
	h.jwtHandler.Sub = strconv.Itoa(int(res.Id))
	h.jwtHandler.Role = "authorized"
	h.jwtHandler.Aud = []string{"pintagol"}
	h.jwtHandler.SiginKey = h.cfg.SiginKey
	h.jwtHandler.Log = h.log
	tokens, err := h.jwtHandler.GenerateAuthJWT()
	accessToken := tokens[0]
	refreshToken := tokens[1]

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error generating token", l.Error(err))
		return
	}
	// keys for update tokens
	ucReq := &pu.RequestForTokens{
		Id:           res.Id,
		RefreshToken: refreshToken,
	}

	// Update tokens
	res, err = h.serviceManager.AdminService().UpdateToken(context.Background(), &pu.RequestForTokens{Id: ucReq.Id, RefreshToken: ucReq.RefreshToken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while updating client token", l.Error(err))
		return
	}

	// Just forresponse
	response := &models.LoginAdmin{
		Id:    res.Id,
		Email: res.Email,
		Name:  res.Name,
	}
	response.AccesToken = accessToken
	response.Refreshtoken = refreshToken
	c.JSON(http.StatusOK, response)
}
