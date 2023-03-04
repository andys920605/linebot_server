package service

type BroadcastMessage struct {
	Message string `json:"message" binding:"required"`
}
