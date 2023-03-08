package comment

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	ErrFetchingComment = errors.New("could not fetch comment by ID")
	ErrUpdatingComment = errors.New("could not update comment")
	ErrNoCommentFound  = errors.New("no comment found")
	ErrDeletingComment = errors.New("could not delete comment")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type CommentStore interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	Ping(context.Context) error
}

type Service struct {
	Store CommentStore
}

func NewService(store CommentStore) *Service {
	return &Service{
		Store: store,
	}
}
func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, ID)
	if err != nil {
		log.Errorf("an error occured fetching the comment: %s", err.Error())
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		log.Errorf("an error occurred adding the comment: %s", err.Error())
	}
	return cmt, nil
}

func (s *Service) UpdateComment(
	ctx context.Context, ID string, newComment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, newComment)
	if err != nil {
		log.Errorf("an error occurred updating the comment: %s", err.Error())
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return s.Store.DeleteComment(ctx, ID)
}

func (s *Service) ReadyCheck(ctx context.Context) error {
	log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}
