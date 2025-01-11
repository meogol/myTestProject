package pc

import (
	"github.com/gin-gonic/gin"
	"meogol/pc-service/common"
	"net/http"
	"strconv"
)

func AddRoutes(r *gin.RouterGroup) {
	router := r.Group("/")
	{
		router.PUT("/pc", updateHandler)
		router.POST("/pc", createHandler)
		router.DELETE("/pc/:id", deleteHandler)
		router.GET("/pc/:id", getHandler)
	}
}

func getHandler(context *gin.Context) {
	pcIdStr := context.Param("id")
	pcId, err := strconv.Atoi(pcIdStr)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponseStr("invalid pc id"))
		return
	}

	pc, err := getPc(pcId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, &pc)
}

func deleteHandler(c *gin.Context) {
	pcIdStr := c.Param("backupId")
	pcId, err := strconv.Atoi(pcIdStr)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponseStr("invalid pc id"))
		return
	}

	err = deletePc(pcId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse())
}

func updateHandler(c *gin.Context) {
	var request *Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	err := updatePc(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse())

}

func createHandler(c *gin.Context) {
	var request *Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	err := createPc(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse())
}
