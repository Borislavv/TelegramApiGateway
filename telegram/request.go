package telegramGateway

// Getting messages struct and interfaces
//
type RequestGetMessagesInterface interface {
	GetOffset() int64
}

type RequestGetMessages struct {
	offset int64
}

// NewRequestGetMessages is a constructor of RequestMessages struct.
func NewRequestGetMessages(offset int64) *RequestGetMessages {
	return &RequestGetMessages{
		offset: offset,
	}
}

// GetOffset is a getter of offset field.
func (reqMsgs *RequestGetMessages) GetOffset() int64 {
	return reqMsgs.offset
}

// Sending messages structs and interfaces
//
type RequestSendMessageInterface interface {
	GetChatId() int64
	GetMessage() string
}

type RequestSendMessage struct {
	chatId  int64
	message string
}

// NewRequestSendMessage is a constructor of RequestSendMessage struct.
func NewRequestSendMessage(chatId int64, message string) *RequestSendMessage {
	return &RequestSendMessage{
		chatId:  chatId,
		message: message,
	}
}

// GetChatId is a getter of chatId field.
func (reqMsg *RequestSendMessage) GetChatId() int64 {
	return reqMsg.chatId
}

// GetMessage is a getter of message field.
func (reqMsg *RequestSendMessage) GetMessage() string {
	return reqMsg.message
}
