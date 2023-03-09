package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/go-playground/validator/v10"

	"github.com/co-codin/service-layer/internal/comment"
)

type CommentService interface {
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func commentFromPostCommentRequest(u PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   u.Slug,
		Author: u.Author,
		Body:   u.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var postCmtReq PostCommentRequest


	if err := json.NewDecoder(r.Body).Decode(&postCmtReq); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(postCmtReq)

	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt := commentFromPostCommentRequest(postCmtReq)

	cmt, err = h.Service.PostComment(r.Context(), cmt)

	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
