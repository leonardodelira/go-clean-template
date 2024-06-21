package translationhdl

import (
	"leonardodelira/go-clean-template/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Msg interface{} `json:"msg"`
}

type httphdl struct {
	service ports.TranslationService
}

func NewTranslationHandler(service ports.TranslationService) *httphdl {
	return &httphdl{
		service: service,
	}
}

func (h *httphdl) DoTranslation(c *gin.Context) {
	input, err := ParseRequest(c)
	if err != nil {
		//todo: instancia do logger
		badRequest(c, err.Error())
		return
	}
	ctx := c.Request.Context()
	result, err := h.service.DoTranslation(ctx, *input)
	if err != nil {
		//todo: log error
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, Response{Data: result})
}

func (h *httphdl) GetTranslations(c *gin.Context) {
	ctx := c.Request.Context()
	t, err := h.service.GetTranslation(ctx)
	if err != nil {
		//todo: log error
		badRequest(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{Data: t})
}

func badRequest(c *gin.Context, msg ...string) {
	messageError := "some error occur, try again after"
	if len(msg) > 0 {
		messageError = msg[0]
	}
	c.JSON(http.StatusBadRequest, map[string]interface{}{"error": messageError})
}
