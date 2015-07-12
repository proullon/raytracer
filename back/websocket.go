package back

import (
    "net/http"
    "github.com/gorilla/websocket"
    "log"
    "container/list"
    "encoding/json"
)

func Raytracer(w http.ResponseWriter, r *http.Request) {
    log.Printf("New connection\n")

    // Upgrade from http request to WebSocket.
    log.Printf("Upgrading to websocket\n")
    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400)
        return
    } else if err != nil {
        log.Printf("Cannot setup WebSocket connection:%s\n", err)
        return
    }

    // Create scene info
    log.Printf("Creating scene info\n")
    objects := list.New()
    objects.PushBack(NewSphere(255, 75))

    eye := NewEye(-170, -50, 0)

    lights := list.New()
    spot := NewLight(-170, -100, 100, 16777215)
    lights.PushBack(spot)

    // Create proxy channel and start write routine
    log.Printf("Creating proxy channel\n")
    channel := make(chan Pixel)
    go broadcastWebSocket(channel, ws)

    // Start raytracing
    log.Printf("start raytracing\n")
    Trace(channel, 900, 900, objects, eye, lights)
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(outChannel chan Pixel, ws *websocket.Conn) {
    i := 0
    buffer := make([]Pixel, 500)

    for {
        select {
        case pixel := <-outChannel:
            if i < 500 {
                buffer[i] = pixel
                i++
            } else {
                i = 0
                data, _ := json.Marshal(buffer)
                ws.WriteMessage(websocket.TextMessage, []byte(data))
            }
            // log.Printf("Pixel received %v\n", pixel)
            // data, _ := json.Marshal(pixel)
            // ws.WriteMessage(websocket.TextMessage, []byte(data))
            break
        }
    }
}

// // Join method handles WebSocket requests for WebSocketController.
// func (this *WebSocketController) Join() {
//     beego.Info("WebSocketControler.Join")

//     // Get username and roomname
//     uname := this.GetString("uname")
//     room := this.GetString("room")
//     if len(uname) == 0 || len(room) == 0 {
//         this.Redirect("/", 302)
//         return
//     }

//     // Upgrade from http request to WebSocket.
//     ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
//     if _, ok := err.(websocket.HandshakeError); ok {
//         http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
//         return
//     } else if err != nil {
//         beego.Error("Cannot setup WebSocket connection:", err)
//         return
//     }

//     // Create new player
//     player := models.NewPlayer(uname)

//     // Join room.
//     inChannel, outPlayer := models.JoinRoom(room, player)

//     go broadcastWebSocket(player.OutChannel, ws)
//     go receiveRoutine(inChannel, ws, player, outPlayer)
// }

// func receiveRoutine(inChannel chan string, ws *websocket.Conn, player *models.Player, outPlayer chan *models.Player) {
//     beego.Info("Starting receive routine")

//     for {
//         _, p, err := ws.ReadMessage()
//         if err != nil {
//             // Time to quit
//             break
//         }

//         inChannel <- string(p)
//     }

//     outPlayer <- player
// }

