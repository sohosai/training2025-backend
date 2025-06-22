package hadlers

import (
	"fmt"
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
				c.IndentedJSON(http.StatusInternalServerError, models.Message{Message: fmt.Sprintf("<%s>", err)})
				return
			}

			postModels := utils.Map(posts, models.IntoPostModel)
			c.IndentedJSON(http.StatusOK, postModels)
		})

	r.POST("/posts",

		func(c *gin.Context) {
			var newPostrequest models.CreatePostRequest
			err := c.BindJSON(&newPostrequest)

			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, models.Message{Message: fmt.Sprintf("<%s>", err)})

				return
			}
			newPost := newPostrequest.Value()
			id, err := repos.PostRepo.CreatePosts(&newPost)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, models.Message{Message: fmt.Sprintf("<%s>", err)})
				return
			}
			c.IndentedJSON(http.StatusOK, id)
		})
	r.GET("/posts/:id",

		func(c *gin.Context) {
			searchid := c.Param("id")
			posts, err := repos.PostRepo.GetPostsById(searchid)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, models.Message{Message: fmt.Sprintf("<%s>", err)})
				return
			}

			postModels := utils.Map(posts, models.IntoPostModel)
			c.IndentedJSON(http.StatusOK, postModels)
		})

}
