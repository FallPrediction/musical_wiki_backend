package request

type Image struct {
	Name    string `json:"name" binding:"required"`
	Image   string `json:"image" binding:"required"`
	ActorId uint32 `json:"actorId" binding:"required,numeric"`
}
