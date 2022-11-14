package service

import "errors"

var (
	ErrChatNotFound = errors.New("chat with such credentials not found")
)
