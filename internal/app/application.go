package app

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"harmancioglue/url-shortener/internal/config"
	"harmancioglue/url-shortener/internal/domain/service"
	urlRepositoryLayer "harmancioglue/url-shortener/internal/infrastructure/repository/mysql"
	"harmancioglue/url-shortener/internal/services"
)

type Application struct {
	UrlService service.UrlService
}

func Init(config *config.Config) (*Application, error) {
	app := &Application{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database: " + err.Error())
	}

	//repositories
	urlRepository := urlRepositoryLayer.NewUrlRepository(db)

	idGenerator, err := services.NewSnowflakeIDGenerator(int64(config.WorkerID))
	if err != nil {
		return nil, errors.New("failed to create ID generator: " + err.Error())
	}

	urlService := services.NewUrlService(urlRepository, idGenerator, config)

	app.UrlService = urlService
	return app, nil
}
