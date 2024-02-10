package request

type Credit struct {
	ActorId   uint32 `json:"actorId" binding:"required,numeric"`
	Time      string `json:"time"`
	Place     string `json:"place"`
	Character string `json:"character" binding:"required"`
	Musical   string `json:"musical" binding:"required"`
}
