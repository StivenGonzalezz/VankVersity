package service

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	muxgo "github.com/muxinc/mux-go/v7"
	"video-service/internal/domain/ports"
)

type VideoService struct {
	client *muxgo.APIClient
	Repo ports.VideoRepository
}

func NewVideoService(Repo ports.VideoRepository) *VideoService {
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(
				os.Getenv("ACCESS_TOKEN_ID"),
				os.Getenv("SECRET_KEY"),
			),
		),
	)
	return &VideoService{client: client, Repo: Repo}
}

func (vs *VideoService) GenerateUploadLink(c *fiber.Ctx) error {
	req := muxgo.CreateUploadRequest{
		NewAssetSettings: muxgo.CreateAssetRequest{
			PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC},
		},
		Timeout:    3600,
		CorsOrigin: "*",
	}

	upload, err := vs.client.DirectUploadsApi.CreateDirectUpload(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("failed to create upload: %v", err)})
	}
	return c.JSON(upload.Data)	
}

func (vs *VideoService) ListVideos(c *fiber.Ctx) error {
	videos, err := vs.client.AssetsApi.ListAssets()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("failed to list videos: %v", err)})
	}
	return c.JSON(videos.Data)
}

func (vs *VideoService) GetVideoByID(id string) (muxgo.AssetResponse, error) {
	return vs.client.AssetsApi.GetAsset(id)
}

func (vs *VideoService) DeleteVideo(id string) error{
	return vs.client.AssetsApi.DeleteAsset(id)
}
