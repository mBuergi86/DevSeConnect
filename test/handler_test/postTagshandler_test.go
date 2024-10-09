package handlertest

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/handler"
	"github.com/stretchr/testify/mock"
)

type MockPostTagsService struct {
	mock.Mock
}

func (m *MockPostTagsService) GetPostTags(ctx context.Context) ([]*entity.PostTags, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.PostTags), args.Error(1)
}

func (m *MockPostTagsService) GetTag(ctx context.Context, tagID uuid.UUID) ([]*entity.PostTags, error) {
	args := m.Called(ctx, tagID)
	return args.Get(0).([]*entity.PostTags), args.Error(1)
}

func (m *MockPostTagsService) Create(ctx context.Context, tag *entity.PostTags, title, tags string) error {
	args := m.Called(ctx, tag, title, tags)
	return args.Error(0)
}

func TestGetPostTags(t *testing.T) {
	e := echo.New()
	mockService := new(MockPostTagsService)
	h := handler.NewPostTagsHandler(mockService)

	postTagID := uuid.New()
	postTag := &entity.PostTags{
		PostID: postTagID,
		TagID:  uuid.New(),
	}

	mockService.On("GetTag", mock.Anything, postTagID).Return(postTag, nil)
}
