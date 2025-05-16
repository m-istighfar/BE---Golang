package ginutils

import (
	"Yuk-Ujian/internal/constant"
	"Yuk-Ujian/internal/dto"
	"Yuk-Ujian/internal/dto/pagedto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// convenience method, errors should be handled implicitly by the middleware
func ResponseOKPlain(c *gin.Context) {
	ResponseOKData(c, nil)
}

func ResponseOKData(c *gin.Context, data interface{}) {
	ResponseOK(c, constant.ResponseMsgSuccess, data)
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	ResponseSuccessJSON(c, http.StatusOK, message, data)
}

func ResponseSuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, dto.SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func ResponseOKPaginated(c *gin.Context, entries interface{}, pageInfo pagedto.PageInfoDto) {
	ResponseOKData(c, dto.PaginatedResponse{
		Entries:  entries,
		PageInfo: pageInfo,
	})
}
