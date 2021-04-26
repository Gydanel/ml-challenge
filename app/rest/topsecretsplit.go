package rest

import (
	"github.com/gin-gonic/gin"
	"ml-challenge/app/api"
	"ml-challenge/app/rest/request"
	"ml-challenge/domain/service"
	"net/http"
)

type TopSecretSplitHandler struct {
	messageService   *service.MessageService
	satelliteService *service.SatelliteService
	decoderService   *service.DecoderService
}

func (h *TopSecretSplitHandler) Init() {
	h.messageService = service.NewMessageService()
	h.satelliteService = service.NewSatelliteService()
}

func (h TopSecretSplitHandler) AddSignal(c *gin.Context) {
	var body request.TopSecretSplitRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	if name := c.Param("satellite_name"); name != "" {
		satellite, err := h.satelliteService.ByName(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error invalid satellite", "error": err})
			c.Abort()
			return
		}
		if _, err := h.messageService.BySatelliteName(name); err != nil {
			h.messageService.Add(body.Distance, body.Message, *satellite)
		} else {
			h.messageService.UpdateBySatelliteName(body.Distance, body.Message, satellite.Name)
		}
		c.JSON(http.StatusOK, body)
		return
	}
}

func (h TopSecretSplitHandler) DecodeMessage(c *gin.Context) {
	var m1, m2, m3 api.MessageApi
	messages := h.messageService.All()
	if len(messages) < 3 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not enough info"})
		return
	}
	m1.Init(*messages[0])
	m2.Init(*messages[1])
	m3.Init(*messages[2])

	d, err := h.decoderService.Decode(m1, m2, m3)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, d)
}