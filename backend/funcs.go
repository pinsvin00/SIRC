package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)


func random_color() string {
	str := "#";

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	for i := 0 ; i < 6; i++ {
		shift := r1.Intn(16);
		if shift >= 8 {
			shift += 8
		}
		str += string((shift + int('0')));
	}
	return str;
}

func get_current_time() string {
	hours, minutes, _ := time.Now().Clock()
	return fmt.Sprintf("%2d:%2d", hours, minutes);
}

func enqueue_all(message Message){
	fmt.Println(message)
	for key, _ := range irc.message_queue {
		irc.message_queue[key] = append(irc.message_queue[key], message)
	}
}

func process_command(command string, request map[string]string) bool {
	splitted := strings.Split(command, " ");
	if len(splitted) == 0 {
		return false;
	}

	com := splitted[0];

	if com == "/quit" {
		message := Message{}
		text := strings.Join(splitted[1:], " ");
		message.text = fmt.Sprintf(QUIT_MESSAGE, irc.guid_nickname[request["guid"]],text);
		message.color = "gray"
		message.sender = "";
		message.time = get_current_time();
		delete(irc.guid_nickname,request["guid"])
		enqueue_all(message)
		return true
	}

	if com == "/colour" {
		guid := request["guid"]
		if len(splitted) <= 1{
			return false;
		}
		irc.guid_color[guid] = splitted[1];
		return true
	}

	if com == "/nick" {
		guid := request["guid"]
		if len(splitted) <= 1{
			return false;
		}
		old_nickname := irc.guid_nickname[guid]
		irc.guid_nickname[guid] = splitted[1];
		message := Message{fmt.Sprintf(NICK_CHANGE, old_nickname, irc.guid_nickname[guid]),"", get_current_time(), "gray" }
		enqueue_all(message)
		return true
	}

	return false;
}

func send_credentials(w http.ResponseWriter, r *http.Request){
	request := load_request(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	name := request["nickname"]
	guid := ksuid.New().String()

	irc.nicknames = append(irc.nicknames, name)
	irc.guid_nickname[guid] = name
	irc.message_queue[guid] = make([]Message, 0)
	irc.guid_color[guid] = random_color()
	fmt.Println(irc.message_queue)
	irc.subscribers++

	response := make(map[string]string)
	response["guid"] = guid
	response["success"] = "true"

	jsonByte, _ := json.Marshal(response)
	w.Write(jsonByte)
}

func message_handler(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	request := load_request(r)
	response := make(map[string]string)
	fmt.Println(irc.guid_color)
	fmt.Println(request["guid"])
	response["success"] = "false"
	if _, ok := irc.guid_nickname[request["guid"]]; ok {
		if !process_command(request["text"], request){
			message := Message{}
			message.sender = irc.guid_nickname[request["guid"]]
			message.text = request["text"]
			message.time = get_current_time()
			message.color = irc.guid_color[request["guid"]]
			response["success"] = "true"
			enqueue_all(message)
		} else {
			response["success"] = "true"
		}
		
	} 


	jsonByte, _ := json.Marshal(response)
	w.Write(jsonByte)
}

func subscribe(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	request := load_request(r)
	fmt.Println(request)
	guid := request["guid"]
	for {
		time.Sleep(time.Millisecond*1000)
		if len(irc.message_queue[guid]) != 0 {
			break;
		}
	}	

	fmt.Println("Messages!")
	fmt.Println(irc.message_queue[guid])

	messages := make([]map[string]string, 0)

	for _,value := range irc.message_queue[guid] {
		dict := make(map[string]string)
		dict["text"] = value.text;
		dict["sender"] = value.sender;
		dict["time"] = value.time;
		dict["color"] = value.color;
		messages = append(messages, dict)
	}
	jsonByte, _ := json.Marshal(messages)
	fmt.Fprintf(os.Stdout, "%s", jsonByte)
	w.Write(jsonByte)
	irc.message_queue[guid] = irc.message_queue[guid][:0]
}