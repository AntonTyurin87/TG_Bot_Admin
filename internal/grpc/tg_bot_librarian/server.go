package tg_bot_librarian

import (
	//"context"
	//"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	//tg_bot_lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/tg_bot_librarian"
)

type Server struct {
	lib.UnimplementedLibrarianServer
	lib.LibrarianClient
}

//func (s *Server)   {
//
//}
//
//func (s *Server) SendMessage(ctx context.Context, in *tg_bot_lib.SendMessageRequest) (*tg_bot_lib.SendMessageResponse, error) {
//
//	res, err := s.SendFile(ctx, &lib.SendFileRequest{
//		Text: in.Text,
//	})
//	if err != nil {
//		fmt.Println("Не получили файл")
//		return nil, err
//	}
//
//	fmt.Printf("Воти содержимое файла - %s", res)
//
//	return &tg_bot_lib.SendMessageResponse{}, nil
//}
