package dao

import (
	"github.com/sarailQAQ/destroy-demo/model"
	"log"
)

func UploadImage(url string) (uint, error) {
	 imageTable := model.ImageTable {
	 	Url: url,
	 }
	err := DB.Model(&imageTable).Create(&imageTable).Error
	if err != nil {
		log.Println("dao-image: ", err)
		return 0, err
	}
	return imageTable.ID, nil
}

func GetImageUrl(id uint)  (string, error){
	var imageTable model.ImageTable
	err := DB.Model(&imageTable).Where("id = ?", id).First(&imageTable).Error
	if err != nil {
		log.Println("dao-image: ", err)
		return "", err
	}
	return imageTable.Url, nil
}

func GetSeveralImages(ids []uint) []string {
	urls := make([]string, 0)
	for _, id := range ids {
		url, _ := GetImageUrl(id)
		urls = append(urls, url)
	}
	return urls
}