package usecase

import (
	"context"
	"go-tinder/domain"
	"log"
	"time"
)

func (u *userUsecase) Like(ctx context.Context, idLiker int64, idLiked int64, isSuper bool) error {
	// Checks if possible Match
	userLiker, err := u.userRepository.GetByID(ctx, idLiker)
	if err != nil {
		log.Println("Error while querying User Liker")
		return err
	}
	likesReceivedByLiker := userLiker.LikesReceived
	for _, likeReceivedByLiker := range likesReceivedByLiker {
		// Match
		if likeReceivedByLiker.IDLiker == idLiked {
			newMatch := domain.Match{IDUser1: idLiker, IDUser2: idLiked, CreatedAt: time.Now().UTC()}
			matchesLiker := append(userLiker.Matches, newMatch)
			userLiker.Matches = matchesLiker
			u.userRepository.Store(ctx, &userLiker)

			userLiked, err := u.userRepository.GetByID(ctx, idLiked)
			if err != nil {
				log.Println("Error while querying User Liked")
				return err
			}
			matchesLiked := append(userLiked.Matches, newMatch)
			userLiked.Matches = matchesLiked
			u.userRepository.Store(ctx, &userLiked)
		}
	}

	// If not yet Match
	userLiked, err := u.userRepository.GetByID(ctx, idLiked)
	if err != nil {
		log.Println(err)
		return err
	}
	newLike := domain.Like{IDLiker: idLiker, IsSuper: isSuper, CreatedAt: time.Now().UTC()}
	updatedLikes := append(userLiked.LikesReceived, newLike)
	userLiked.LikesReceived = updatedLikes
	err = u.userRepository.Store(ctx, &userLiked)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *userUsecase) Nope(ctx context.Context, idNoper int64, idNoped int64) error {
	// Checks if Noper had been liked
	userNoper, err := u.userRepository.GetByID(ctx, idNoper)
	if err != nil {
		log.Println("Error while querying User Noper")
		return err
	}
	likesReceivedByNoper := userNoper.LikesReceived
	for i, likeReceivedByNoper := range likesReceivedByNoper {
		// Had been liked
		if likeReceivedByNoper.IDLiker == idNoped {
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
