package types

import "time"

type ChatUser struct {
	Id string `json:"id"`
}

type ChatChannel struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	ActiveUsers      []ChatUser
	BroadcastChannel chan string
}

type Message struct {
	Id         string    `json:"id"`
	AuthorId   string    `json:"authorId"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	ClientUUID string    `json:"clientUUID"`
	RoomId     string    `json:"roomId"`
	Author     struct {
		Id            string      `json:"id"`
		Name          string      `json:"name"`
		Email         string      `json:"email"`
		EmailVerified interface{} `json:"emailVerified"`
		Image         string      `json:"image"`
		ChannelId     interface{} `json:"channelId"`
	} `json:"author"`
	Room struct {
		Id        string `json:"id"`
		OwnerId   string `json:"ownerId"`
		IsPrivate bool   `json:"isPrivate"`
		Name      string `json:"name"`
	} `json:"room"`
}

type ChatMessage struct {
	Id      int64       `json:"id"`
	Owner   ChatUser    `json:"owner"`
	Channel ChatChannel `json:"channel"`
}

type WsServer struct {
	RegisterChannel   chan *ChatUser
	UnregisterChannel chan *ChatUser
}

type JoinRoomRequest struct {
	roomName string `json:"roomName"`
	jwt      string `json:jwt`
}
