package chat

import (
	"errors"
	"github.com/NolikTop/watcher/pkg/config"
)

func NewMethod(config *config.NotificationMethodConfig) (Chat, error) {
	var notificationMethod Chat

	protocol := config.Method
	switch protocol {
	case "vk":
		notificationMethod = &Vk{}
	default:
		return nil, errUnknownMethod(protocol)
	}

	err := notificationMethod.Init(config.Name, config.Data)
	if err != nil {
		return nil, err
	}

	return notificationMethod, nil
}

func errUnknownMethod(serverProtocol string) error {
	return errors.New("Unknown server \"" + serverProtocol + "\"")
}