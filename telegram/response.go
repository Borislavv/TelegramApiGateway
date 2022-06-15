package telegramGateway

import (
	"errors"
	"fmt"
)

// Getting messages struct and interfaces
//
type ResponseGetMessagesInterface interface {
	GetMessages() []ResponseGetMessage
}

type ResponseGetMessages struct {
	Messages []ResponseGetMessage `json:"result"`
}

type ResponseGetMessage struct {
	QueueId int64 `json:"update_id"`
	Data    struct {
		Chat struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int64  `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

// NewResponseGetMessages is a constructor of ResponseGetMessages struct.
func NewResponseGetMessages() *ResponseGetMessages {
	return &ResponseGetMessages{}
}

// GetMessages is a getter of ResponseGetMessage slice.
func (respMsgs *ResponseGetMessages) GetMessages() []ResponseGetMessage {
	return respMsgs.Messages
}

// GetQueueId is a getter of QueueId field.
func (respMsg ResponseGetMessage) GetQueueId() int64 {
	return respMsg.QueueId
}

// GetChatId is a getter of ChatId field.
func (respMsg ResponseGetMessage) GetChatId() int64 {
	return respMsg.Data.Chat.ID
}

// GetTitle is a getter of FirstName field.
func (respMsg ResponseGetMessage) GetFirstName() string {
	return respMsg.Data.Chat.FirstName
}

// GetTitle is a getter of FirstName field.
func (respMsg ResponseGetMessage) GetLastName() string {
	return respMsg.Data.Chat.LastName
}

// GetUsername is a getter of Username field.
func (respMsg ResponseGetMessage) GetUsername() string {
	return respMsg.Data.Chat.Username
}

// GetUsername is a getter of Chat.Type field.
func (respMsg ResponseGetMessage) GetChatType() string {
	return respMsg.Data.Chat.Type
}

// GetDate is a getter of Date field.
func (respMsg ResponseGetMessage) GetDate() int64 {
	return respMsg.Data.Date
}

// GetText is a getter of Text field.
func (respMsg ResponseGetMessage) GetText() string {
	return respMsg.Data.Text
}

// Sending messages structs and interfaces
//
type ResponseSendMessageInterface interface {
	IsOK() bool
	ToError() error
}

type ResponseSendMessage struct {
	Status      bool `json:"ok"`
	RawResponse []byte
}

// NewResponseSendMessage is a constructor of ResponseSendMessage struct.
func NewResponseSendMessage() *ResponseSendMessage {
	return &ResponseSendMessage{}
}

// IsOK is a getter of Status field.
func (respMsg *ResponseSendMessage) IsOK() bool {
	return respMsg.Status
}

// ToError method will return an error with received body into (used when status is not OK).
func (respMsg *ResponseSendMessage) ToError() error {
	return errors.New("received status not OK: " + fmt.Sprintf("%+v\n", string(respMsg.RawResponse)))
}
