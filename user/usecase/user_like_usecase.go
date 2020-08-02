package usecase

import (
	"context"
	"go-tinder/domain"
	"log"
	"time"
)

func (u *userUsecase) Like(ctx context.Context, userLiker domain.User, userLiked domain.User, isSuper bool) error {
	// Checks if possible Match
	likesReceivedByLiker := userLiker.LikesReceived
	for _, likeReceivedByLiker := range likesReceivedByLiker {
		// Match
		if likeReceivedByLiker.IDLiker == userLiked.ID {
			newMatchLiker := domain.Match{IDUserMatched: userLiked.ID, CreatedAt: time.Now().UTC()}
			matchesLiker := append(userLiker.Matches, newMatchLiker)
			userLiker.Matches = matchesLiker
			err := u.userRepository.Store(ctx, &userLiker)
			if err != nil {
				log.Println("Error while storing new match")
				return err
			}

			newMatchLiked := domain.Match{IDUserMatched: userLiker.ID, CreatedAt: time.Now().UTC()}
			matchesLiked := append(userLiked.Matches, newMatchLiked)
			userLiked.Matches = matchesLiked
			err = u.userRepository.Store(ctx, &userLiked)
			if err != nil {
				log.Println("Error while storing new match")
				return err
			}
		}
	}

	// If not yet Match
	newLike := domain.Like{IDLiker: userLiker.ID, IsSuper: isSuper, CreatedAt: time.Now().UTC()}
	updatedLikes := append(userLiked.LikesReceived, newLike)
	userLiked.LikesReceived = updatedLikes
	err := u.userRepository.Store(ctx, &userLiked)
	if err != nil {
		log.Println("Error while storing new like")
		return err
	}
	return nil
}

func (u *userUsecase) Nope(ctx context.Context, userNoper domain.User, userNoped domain.User) error {
	// Checks if Noper had been liked
	likesReceivedByNoper := userNoper.LikesReceived
	for i, likeReceivedByNoper := range likesReceivedByNoper {
		// Had been liked
		if likeReceivedByNoper.IDLiker == userNoped.ID {
			// Remove Noped from Noper list of received likes
			likesReceivedByNoper[i] = likesReceivedByNoper[len(likesReceivedByNoper)-1]
			userNoper.LikesReceived = likesReceivedByNoper[:len(likesReceivedByNoper)-1]
			err := u.userRepository.Store(ctx, &userNoper)
			if err != nil {
				log.Println("Error while storing User Noper")
				return err
			}
		}
	}
	return nil
}
