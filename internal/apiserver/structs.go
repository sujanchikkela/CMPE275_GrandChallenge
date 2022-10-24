package apiserver

type PostParams struct {
	Name     string `json:"name"`
	RoomName string `json:"roomname"`
	Message  string `json:"message"`
}
