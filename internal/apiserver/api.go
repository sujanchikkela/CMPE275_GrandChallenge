package apiserver

import ( "github.com/go-chi/render" "grandChallenge1/grpcchat123/conf" "grandChallenge1/grpcchat123/internal/chatserver/client"
	"grandChallenge1/grpcchat123/internal/chatserver/client1" "encoding/json" "fmt" "io/ioutil" "net/http" "strings"
        )
type Api struct { Cnf *conf.Conf }
func (api *Api) GetMessage(w http.ResponseWriter, r *http.Request) 
{
	data, err := ioutil.ReadFile(api.Cnf.ChatLog)
	if err != nil 
	{ w.WriteHeader(http.StatusInternalServerError) w.Write([]byte("Internal server error")) return }
	line_arr := strings.Split(string(data), "\n")
	render.JSON(w, r, line_arr)
}
func (api *Api) PostMessage(w http.ResponseWriter, r *http.Request) { name := "Satvik" room_name := "default" blocked_name := "NA"
	params := &PostParams{} err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil { w.WriteHeader(http.StatusInternalServerError) w.Write([]byte("Internal server error")) return }
	if params.Name == "" {	params.Name = name }
	if params.RoomName == "" {
		params.RoomName = room_name
	}
	client := &client.Client{
		Name:        params.Name,
		RoomName:    params.RoomName,
		BlockedName: blocked_name,
	}
	fmt.Println(client.Name, client.RoomName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
}
