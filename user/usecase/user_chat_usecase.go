package usecase

import (
	"go-tinder/domain"
	"context"
	"time"
	"log"
)

func (u *userUsecase) sendChatMessage(ctx context.Context, userSender domain.User, userReceiver domain.User, text string) error {
	chatMessage := domain.ChatMessage{IDSender: userSender.ID, IDReceiver: userReceiver.ID, Text: text, CreatedAt: time.Now().UTC()}
	for _, matchWithReceiver := range userSender.Matches {
		if matchWithReceiver.IDUserMatched == userReceiver.ID {
			matchWithReceiver.ChatMessages = append(matchWithReceiver.ChatMessages, chatMessage)
			err := u.userRepository.Store(ctx, &userSender)
			if err != nil {
				log.Println("Error while storing new chat message for User Sender")
				return err
			}
		}
	}
	for _, matchWithSender := range userReceiver.Matches {
		if matchWithSender.IDUserMatched == userSender.ID {
			matchWithSender.ChatMessages = append(matchWithSender.ChatMessages, chatMessage)
			err := u.userRepository.Store(ctx, &userReceiver)
			if err != nil {
				log.Println("Error while storing new chat message for User Receiver")
				return err
			}
		}
	}
	return nil
}

func (u *userUsecase) likeMessage(ctx context.Context, userLiker domain.User, userLiked domain.User, chatMessageLiked domain.ChatMessage, isLike bool) error {
	for _, matchWithLiked := range userLiker.Matches {
		if matchWithLiked.IDUserMatched == userLiked.ID {
			for _, chatMessage := range matchWithLiked.ChatMessages {
				if chatMessage.ID == chatMessageLiked.ID {
					chatMessage.Like = isLike
					err := u.userRepository.Store(ctx, &userLiker)
					if err != nil {
						log.Println("Error while storing like for Chat Message")
						return err
					}
				}
			}
		}
	}
	for _, matchWithLiker := range userLiked.Matches {
		if matchWithLiker.IDUserMatched == userLiker.ID {
			for _, chatMessage := range matchWithLiker.ChatMessages {
				if chatMessage.ID == chatMessageLiked.ID {
					chatMessage.Like = isLike
					err := u.userRepository.Store(ctx, &userLiked)
					if err != nil {
						log.Println("Error while storing like for Chat Message")
						return err
					}
				}
			}
		}
	}
	return nil
}