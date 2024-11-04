package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Message struct {
	Type      string  `json:"type"`
	Username  string  `json:"username,omitempty"`
	Message   string  `json:"message,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Color     string  `json:"color,omitempty"`
	Thickness string `json:"thickness,omitempty"`
	UserId   string `json:"userId,omitempty"`
	UserList []User `json:"userList,omitempty"`
	RoomId   string `json:"roomId,omitempty"`
}

type User struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

var (
	clients   = make(map[*websocket.Conn]string)
	usernames = make(map[string]string)
	mutex     sync.Mutex
)

type Room struct {
	ID    string
	Users map[*websocket.Conn]string
	Mutex sync.Mutex
}

var (
	rooms = make(map[string]*Room)
)
func createRoom(roomId string) {
    rooms[roomId] = &Room{
        ID:    roomId,
        Users: make(map[*websocket.Conn]string),
    }
}

func joinRoom(ws *websocket.Conn, roomId string, username string) error {
    room, exists := rooms[roomId]
    if !exists {
        return fmt.Errorf("room does not exist")
    }

    room.Mutex.Lock()
    defer room.Mutex.Unlock()
    room.Users[ws] = username

    return nil
}

func broadcastRoomMessage(roomId string, msg Message) {
    room, exists := rooms[roomId]
    if !exists {
        return
    }

    room.Mutex.Lock()
    defer room.Mutex.Unlock()

    for client := range room.Users {
        _ = websocket.JSON.Send(client, msg)
    }
}


func generateUniqueUserID() string {
	return uuid.New().String()
}

func handleClient(ws *websocket.Conn) {
	defer func() {
		err := ws.Close()
		if err != nil {
			log.Printf("Erreur lors de la fermeture de la connexion WebSocket : %v\n", err)
		}
	}()

	var userId string
	var username string

	mutex.Lock()
	clients[ws] = ""
	mutex.Unlock()

	var msg Message

	for {
		var msg Message
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			log.Printf("Erreur lors de la réception du message : %v\n", err)
			break
		}

		if msg.Type == "setUsername" {
			userId = generateUniqueUserID()
			username = msg.Username

			mutex.Lock()
			clients[ws] = userId
			usernames[userId] = username
			mutex.Unlock()

			response := Message{
				Type:   "userId",
				UserId: userId,
			}
			if err := websocket.JSON.Send(ws, response); err != nil {
				log.Printf("Erreur lors de l'envoi du UserId : %v\n", err)
			}

			if err := broadcastUserList(); err != nil {
				log.Printf("Erreur lors de la diffusion de la liste des utilisateurs : %v\n", err)
			}
			fmt.Printf("Utilisateur connecté: %s avec l'ID %s\n", username, userId)
		} else if msg.Type == "message" {
			fmt.Printf("Message reçu de %s: %s\n", username, msg.Message)
			if err := broadcastMessage(msg); err != nil {
				log.Printf("Erreur lors de la diffusion du message : %v\n", err)
			}
		}

	}

	mutex.Lock()
	fmt.Printf("Utilisateur déconnecté: %s avec l'ID %s\n", username, userId)
	delete(clients, ws)
	delete(usernames, userId)
	mutex.Unlock()

	if err := broadcastUserList(); err != nil {
		log.Printf("Erreur lors de la diffusion de la liste après déconnexion : %v\n", err)
	}
}

func broadcastMessage(msg Message) error {
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		if err := websocket.JSON.Send(client, msg); err != nil {
			log.Printf("Erreur lors de l'envoi du message au client : %v\n", err)
			return err
		}
	}
	return nil
}

func broadcastUserList() error {
	mutex.Lock()
	defer mutex.Unlock()

	userList := make([]User, 0, len(usernames))
	for id, name := range usernames {
		userList = append(userList, User{UserId: id, Username: name})
	}
	fmt.Println("Liste des utilisateurs connectés :", userList)

	for client := range clients {
		if err := websocket.JSON.Send(client, Message{Type: "userList", UserList: userList}); err != nil {
			log.Printf("Erreur lors de l'envoi de la liste des utilisateurs : %v\n", err)
			return err
		}
	}
	return nil
}

func main() {
	http.Handle("/websocket", websocket.Handler(handleClient))

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	fmt.Println("Serveur démarré sur : http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v\n", err)
	}
}
