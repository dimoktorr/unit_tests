package api

import (
	"context"
	v1 "github.com/dimoktorr/unit_tests/integration/pkg/api/v1"
	"log"
)

type Server struct {
	v1.UnimplementedExampleServiceServer
}

func (s *Server) GetExample(ctx context.Context, in *v1.GetRequest) (*v1.GetResponse, error) {
	_ = in.GetId()

	//TODO: implement logic here

	//TODO: что-то пишем в кафку используем producer

	log.Println("GetExample started")
	defer log.Println("GetExample finished")

	return &v1.GetResponse{
		Examples: []*v1.Example{
			{
				FirstName:   "first_name",
				LastName:    "last_name",
				Description: "description",
			},
		},
	}, nil
}
