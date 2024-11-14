package integration

import (
	"context"
	v1 "github.com/dimoktorr/unit_tests/integration/pkg/api/v1"
	"github.com/dimoktorr/unit_tests/integration/tests/util"
)

func (s *IntegrationServerSuite) TestGetExample() {
	//Arrange
	want := &v1.GetResponse{
		Examples: []*v1.Example{
			{
				FirstName:   "first_name",
				LastName:    "last_name",
				Description: "description",
			},
		},
	}

	req := &v1.GetRequest{
		Id: 826452,
	}

	//Act
	got, err := s.exampleGrpcClient.GetExample(context.Background(), req)
	//это то, о чем я говорил, s.Require() остановит выполнение теста если err != nil
	s.Require().Nil(err)

	//Assert
	util.NewProtoEqual(s, want, got)

	//TODO: читаем из кафки сообщение консюмером которое отправил продюсер
}
