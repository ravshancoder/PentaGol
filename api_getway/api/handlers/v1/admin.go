package v1

import (
	"context"
	"net/http"
	"strconv"

	pu "github.com/PentaGol/api_getway/genproto/admin"
	l "github.com/PentaGol/api_getway/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// @Summary get user by id
// @Description This api gets a user by id
// @Tags Admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.Admin
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/admin/{id} [get]
func (h *handlerV1) GetAdminById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.serviceManager.AdminService().GetAdminById(context.Background(), &pu.IdRequest{Id: int64(id)})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if status.Code(err) == codes.NotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user by id: ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
