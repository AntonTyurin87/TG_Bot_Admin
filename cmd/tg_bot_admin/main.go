package main

import (
	"TG_Bot_Admin/internal/pkg/domain/presenter"
	"TG_Bot_Admin/internal/pkg/service/admin"
	"TG_Bot_Admin/internal/pkg/service/librarian"
	"TG_Bot_Admin/internal/pkg/service/repository"
	"TG_Bot_Admin/internal/pkg/service/storage"
	telegram "TG_Bot_Admin/internal/pkg/service/telegram"
	telegram2 "TG_Bot_Admin/internal/pkg/service/telegram/handlers"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port                    = ":50053"
	serviceLibrarianAddress = "localhost:50052" //"librarian-app:50052"
)

func main() {
	// Основной контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// подключение gRPC клиентов
	librarianConnect, err := grpc.NewClient(serviceLibrarianAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Не подключился к сервису Librarian")
	}
	defer librarianConnect.Close() //TODO обработка ошибки

	librarianClient := lib.NewLibrarianClient(librarianConnect)
	librarianImpl := librarian.NewLibrarian(librarianClient)

	// подключение БД сервиса
	db, err := storage.InitDB()
	if err != nil {
		fmt.Println("Подключиться к БД RC_Admin_bot.db не удалось - ", err)
	}
	storageImpl := storage.NewStorage(db)

	repositoryImpl := repository.NewRepository(storageImpl)

	presenterImpl := presenter.New()

	// формирование сервиса
	adminService := admin.NewAdminService(presenterImpl, librarianImpl, repositoryImpl)

	// Создание контроллера/обработчика
	handler := telegram2.NewHandler(adminService, presenterImpl)

	// Создание Telegram бота
	telegramBot, err := telegram.CreateTelegramBot(handler)
	fmt.Errorf("создание TG_Bot_Admin не состоялось! - %w", err) //TODO обработка ошибки
	if err != nil {
		panic(err) //TODO а может как-то иначе надо
	}

	// Запуск бота
	//go func() {
	telegramBot.Start(ctx)
	//}()
}
