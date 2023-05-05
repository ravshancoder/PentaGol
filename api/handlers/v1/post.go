package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PentaGol/api/handlers/models"
	pu "github.com/PentaGol/genproto/post"
	l "github.com/PentaGol/pkg/logger"
	"github.com/PentaGol/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary create post
// @Description This api creates a post
// @Tags Post
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.PostRequest true "CreatePost"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/post [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        models.PostRequest
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

	response, err := h.strg.Post().CreatePost(&pu.PostRequest{
		Title:       body.Title,
		Description: body.Description,
		ImgUrl:      body.ImgUrl,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary update post
// @Description This api updates a post
// @Tags Post
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.UpdatePostRequest true "UpdatePost"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandartErrorModel
// @Failure 404 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/post/{id} [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        models.UpdatePostRequest
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

	err = h.strg.Post().UpdatePost(&pu.UpdatePostRequest{
		Id:          body.Id,
		Title:       body.Title,
		Description: body.Description,
		ImgUrl:      body.ImgUrl,
	})

	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
			h.log.Error("Post not found", l.Error(err))
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update Post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, err)
}

// @Summary get post by id
// @Description This api gets a post by id
// @Tags Post
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/post/{id} [get]
func (h *handlerV1) GetPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.strg.Post().GetPostById(&pu.IdRequest{Id: int64(id)})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if status.Code(err) == codes.NotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post by id: ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary get all posts
// @Description This api gets all posts
// @Tags Post
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {object} []models.Posts
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/posts [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errstr := utils.ParseQueryParams(queryParams)
	if errstr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errstr[0],
		})
		h.log.Error("Failed to get all posts: " + errstr[0])
		return
	}
	fmt.Println(params.Limit, params.Page)
	response, err := h.strg.Post().GetAllPosts(&pu.AllPostRequest{
		Limit: params.Limit,
		Page:  params.Page,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get all posts: ", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @Summary delete post
// @Description This api deletes a post
// @Tags Post
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/post/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.strg.Post().DeletePost(&pu.IdRequest{
		Id: int64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
