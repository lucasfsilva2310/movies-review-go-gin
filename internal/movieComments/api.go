package movieComments

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasfsilva2310/movies-review/internal/middlewares"
	"github.com/lucasfsilva2310/movies-review/internal/utils"
)

func RegisterMovieCommentsRoutes(apiConnection *gin.Engine, service *MovieCommentService) {
	movieCommentsURL := apiConnection.Group("/movie-comments")

	movieCommentsURL.POST("/", middlewares.AdminMiddleware(), func(ctx *gin.Context) {
		var movieComment MovieComment

		if err := ctx.ShouldBindJSON(&movieComment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := service.Create(movieComment)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	movieCommentsURL.GET("/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		movieComments, err := service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movieComments)
	})

	movieCommentsURL.GET("/user/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		movieComments, err := service.GetAllByUserID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movieComments)
	})

	movieCommentsURL.GET("/movie/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		movieComments, err := service.GetAllByMovieID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movieComments)
	})

	movieCommentsURL.PATCH("/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		}

		var movieComment MovieCommentUpdate

		if err := ctx.ShouldBindJSON(&movieComment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body"})
		}

		username, errorUsername := utils.GetUsernameFromContext(ctx)

		if errorUsername != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errorUsername.Error()})
		}

		err := service.UpdateMovieComment(id, username, movieComment)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	movieCommentsURL.DELETE("/:id", middlewares.AdminMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		username, errorUsername := utils.GetUsernameFromContext(ctx)

		if errorUsername != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errorUsername.Error()})
		}

		err := service.Delete(id, username)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})
}
