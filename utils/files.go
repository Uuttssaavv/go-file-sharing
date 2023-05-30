package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadFile(file multipart.File, fileType string) (*uploader.UploadResult, error) {
	cld, _ := cloudinary.NewFromURL(GodotEnv("CLOUDINARY_URL"))

	randString := GenerateRandomString(6)
	fileName := fmt.Sprintf("uploads/%s", randString)
	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: fileName,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetFileUrl(fileId string) string {
	cld, _ := cloudinary.NewFromURL(GodotEnv("CLOUDINARY_URL"))

	result, err := cld.Admin.Asset(context.Background(), admin.AssetParams{
		PublicID: fileId,
	})
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return result.URL
}
