package request

type TopSecretSplitRequest struct {
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}
