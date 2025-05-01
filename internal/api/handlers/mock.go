package handlers

import "context"

func NewMockShortLinkRepo() *MockShortLinkRepo {
	return &MockShortLinkRepo{}
}

type MockShortLinkRepo struct {
	saveURL func(context.Context, string) (string, error)
	getURL  func(context.Context, string) (string, error)
}

func (m *MockShortLinkRepo) SaveURL(ctx context.Context, originalURL string) (string, error) {
	return m.saveURL(ctx, originalURL)
}

func (m *MockShortLinkRepo) GetURL(ctx context.Context, key string) (string, error) {
	return m.getURL(ctx, key)
}
