package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ml-challenge/app/api"
	"ml-challenge/app/rest/request"
	"ml-challenge/domain/service"
	"net/http"
)

type TopSecretHandler struct {
	decoderService   *service.DecoderService
	satelliteService *service.SatelliteService
}

func (h *TopSecretHandler) Init() {
	h.decoderService = service.NewDecoderService()
	h.satelliteService = service.NewSatelliteService()
}

func (h TopSecretHandler) DecodeMessage(c *gin.Context) {
	var body request.TopSecretRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	if len(body.Satellites) != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "should be 3 satellites"})
		return
	}
	messages := [3]api.MessageApi{}
	for i, m := range body.Satellites {
		s, err := h.satelliteService.ByName(m.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid satellite name %s", m.Name)})
			return
		}
		if m.Distance <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s distance should be positive", m.Name)})
			return
		}
		if m.Message == nil || len(m.Message) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s message should be a list of strings", m.Name)})
			return
		}
		messages[i] = api.MessageApi{
			Receiver: api.SatelliteApi{
				Name: s.Name,
				PosX: s.PosX,
				PosY: s.PosY,
			},
			Distance: m.Distance,
			Message:  m.Message,
		}
	}
	res, err := h.decoderService.Decode(messages[0], messages[1], messages[2])
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
