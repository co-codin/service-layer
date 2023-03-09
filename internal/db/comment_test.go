// +build integration

package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/co-codin/service-layer/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {

	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_TABLE", "service_layer")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("SSL_MODE", "disable")

	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()

		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})

		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)

		fmt.Println("testing the creation of comments")
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		cmt.Slug = "new-slug"
		cmt, err = db.UpdateComment(context.Background(), cmt.ID, cmt)
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "new-slug", newCmt.Slug)
	})

	t.Run("test get comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)

	})
}
