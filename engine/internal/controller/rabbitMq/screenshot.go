package rabbitmq

import (
	"context"
	"engine/internal/configs"
	"engine/internal/service/usecase"
	"engine/pkg/browser"
	"engine/pkg/firebase"
	rabbitmqtool "engine/pkg/rabbitMq"
	"engine/pkg/utils"
	"log"
)

type RabbitMq struct {
	cfg configs.Config
	uc usecase.ScreenshotUsecase
}

func NewRabbitMq(
	cfg configs.Config,
	uc usecase.ScreenshotUsecase,
) {
	
	rq:=RabbitMq{
		cfg:cfg,
		uc:uc,

	}
	rq.startConsuming()
}

func(rq *RabbitMq )startConsuming() {
	queueName := "violation"

	// Listen for messages in the background
	go rabbitmqtool.Consume(queueName, func(message string) {
		log.Printf("Received message: %s", message)
		Id,url:=utils.SplitBeforeAfterFirstSlash(message)
		screenshot,err:=CaptureScreenshot(rq.cfg,url)
		if err!=nil{
			log.Println(err)
		}
		fire:=firebase.NewFirebaseStore(context.Background(),&rq.cfg)
		path,err:=fire.UploadFile(context.Background(),Id,screenshot)
		if err!=nil{
			log.Println(err)
		}
		err=rq.uc.CreateScreenshot(context.Background(),Id,path)
		if err!=nil{
			log.Println(err)
		}
	})
}

func CaptureScreenshot(cfg configs.Config, url string) ([]byte, error) {
	browser := browser.InitializeChrome(cfg)
	screenshot, err := browser.CaptureScreenshot(url)
	if err != nil {
		return nil, err
	}
	return screenshot, nil
}