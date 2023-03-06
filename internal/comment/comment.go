package comment

import (
	"context"
	"fmt"
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
