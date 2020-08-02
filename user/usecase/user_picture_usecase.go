package usecase

import (
	"go-tinder/domain"
	"context"
	"image"
	"time"
	"log"
)

func (u *userUsecase) uploadPicture(ctx context.Context, user domain.User, image image.Image) error {
	user.Profile.Pictures = append(user.Profile.Pictures, domain.Picture{Image: image, CreatedAt: time.Now().UTC()})
	err := u.userRepository.Store(ctx, &user)
	if err != nil {
		log.Println("Error while storing new picture")
		return err
	}
	return nil
}

func (u *userUsecase) deletePicture(ctx context.Context, user domain.User, picture domain.Picture) error {
	for i, pictureIteration := range user.Profile.Pictures {
		if pictureIteration.ID == picture.ID {
			user.Profile.Pictures[i] = user.Profile.Pictures[len(user.Profile.Pictures)-1]
			user.Profile.Pictures = user.Profile.Pictures[:len(user.Profile.Pictures)-1]
			err := u.userRepository.Store(ctx, &user)
			if err != nil {
				log.Println("Error while deleting new picture")
				return err
			}
		}
	}
	return nil
}