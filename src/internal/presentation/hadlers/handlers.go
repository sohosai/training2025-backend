package hadlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/domain"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/presentation/hadlers/models"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/utils"
)

func Handle(r *gin.Engine, repos domain.Repositories) {
	r.GET("/health", func(c *gin.Context) {
		message := models.Message{Message: "Hello, World!"}
		c.IndentedJSON(http.StatusOK, message)
	})

	r.GET("/posts",
		func(c *gin.Context) {
			posts, err := repos.PostRepo.GetPosts()
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, models.Message{Message: "Internal Error"})
				return
			}

			postModels := utils.Map(posts, models.IntoPostModel)
			c.IndentedJSON(http.StatusOK, postModels)
		})

	r.POST("/posts", func(c *gin.Context) {
		var req models.CreatePostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdPost, err := repos.PostRepo.CreatePost(domain.CreatePostArgs{
			Text:      req.Text,
			CreatedBy: req.CreatedBy,
		})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, models.Message{Message: "Internal Error"})
			return
		}

		c.IndentedJSON(http.StatusCreated, models.CreatePostResponse{Id: createdPost.Id})
	})
}