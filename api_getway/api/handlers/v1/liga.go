package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PentaGol/api_getway/api/handlers/models"
	pu "github.com/PentaGol/api_getway/genproto/liga"
	l "github.com/PentaGol/api_getway/pkg/logger"
	"github.com/PentaGol/api_getway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// liga
// --------------------------------------------------------
// @Summary create liga
// @Description This api creates a liga
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.LigaRequest true "CreateLiga"
// @Success 200 {object} models.LigaResponse
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/liga [post]
func (h *handlerV1) CreateLiga(c *gin.Context) {

	fmt.Println(c.GetHeader("Authorization"))
	var (
		body        models.LigaRequest
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json: ", l.Error(err))
		return
	}

	response, err := h.serviceManager.LigaService().CreateLiga(context.Background(), &pu.LigaRequest{
		Name: body.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create liga", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary get liga by id
// @Description This api gets a Liga by id
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.LigaResponse
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/liga/{id} [get]
func (h *handlerV1) GetLigaById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.serviceManager.LigaService().GetLigaById(context.Background(), &pu.IdRequest{Id: int64(id)})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if status.Code(err) == codes.NotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get liga by id: ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary get all ligas
// @Description This api gets all ligas
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {object} []models.Ligas
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/ligas [get]
func (h *handlerV1) GetAllLigas(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errstr := utils.ParseQueryParams(queryParams)
	if errstr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errstr[0],
		})
		h.log.Error("Failed to get all ligas: " + errstr[0])
		return
	}
	response, err := h.serviceManager.LigaService().GetAllLigas(context.Background(), &pu.AllLigaRequest{
		Limit: params.Limit,
		Page:  params.Page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get all ligas: ", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @Summary delete liga
// @Description This api deletes a liga
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} models.LigaResponse
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/liga/{id} [delete]
func (h *handlerV1) DeleteLiga(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.serviceManager.LigaService().DeleteLiga(context.Background(), &pu.IdRequest{
		Id: int64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete liga", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// game
// -----------------------------------------------
// @Summary create game
// @Description This api creates a game
// @Tags Game
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.GameRequest true "CreateGame"
// @Success 200 {object} models.GameResponse
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/game [post]
func (h *handlerV1) CreateGame(c *gin.Context) {

	fmt.Println(c.GetHeader("Authorization"))
	var (
		body        models.GameRequest
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json: ", l.Error(err))
		return
	}

	response, err := h.serviceManager.LigaService().CreateGame(context.Background(), &pu.GameRequest{
		Time: body.Time,
		Condtion: body.Condtion,
		FirstTeamId: body.FirstTeamId,
		SecondTeamId: body.SecondTeamId,
		ResultFirstTeam: body.ResultFirstTeam,
		ResultSecondTeam: body.ResultSecondTeam,
		FirstTeamPoint: body.FirstTeamPoint,
		SecondTeamPoint: body.SecondTeamPoint,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create liga", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary get liga by id
// @Description This api gets a Liga by id
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.LigaResponse
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/liga/{id} [get]
func (h *handlerV1) GetGameById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.serviceManager.LigaService().GetLigaById(context.Background(), &pu.IdRequest{Id: int64(id)})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if status.Code(err) == codes.NotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get liga by id: ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary get all ligas
// @Description This api gets all ligas
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {object} []models.Ligas
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/ligas [get]
func (h *handlerV1) GetAllGames(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errstr := utils.ParseQueryParams(queryParams)
	if errstr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errstr[0],
		})
		h.log.Error("Failed to get all ligas: " + errstr[0])
		return
	}
	response, err := h.serviceManager.LigaService().GetAllLigas(context.Background(), &pu.AllLigaRequest{
		Limit: params.Limit,
		Page:  params.Page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get all ligas: ", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @Summary delete liga
// @Description This api deletes a liga
// @Tags Liga
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} models.LigaResponse
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/liga/{id} [delete]
func (h *handlerV1) DeleteGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.serviceManager.LigaService().DeleteLiga(context.Background(), &pu.IdRequest{
		Id: int64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete liga", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
