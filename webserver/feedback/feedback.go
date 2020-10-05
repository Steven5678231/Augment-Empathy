package feedback

type EmotionMsg struct {
	UserID  string      `json:"userID"`
	RoomID  string      `json:"roomID"`
	Type 	string 		`json:"type"`
	Emotion []float32	`json:"value"`
}

type EmotionsRepo struct {
	Emotions []EmotionMsg
}

func New() *EmotionsRepo {
	return &EmotionsRepo{
		Emotions: []EmotionMsg{},
	}
}

func (r *EmotionsRepo) Add(emotion EmotionMsg) {
	r.Emotions = append(r.Emotions, emotion)
}

func (r *EmotionsRepo) GetAll() []EmotionMsg {
	return r.Emotions
}
