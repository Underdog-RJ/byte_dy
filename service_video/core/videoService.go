package core

import (
	"context"
	"fmt"
	"service_video/services"
)

func (s *VideoService) UploadVideo(ctx context.Context, request *services.VideoRequest, response *services.VideoResponse) error {
	response.Code = 200
	fmt.Println(request.Data)
	fmt.Println(request.Title)
	return nil
	//TODO implement me
	//panic("implement me")
}
