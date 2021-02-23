package storage

import "sync"

type TokenStore struct {
	tokens map[string]string
	mu     sync.Mutex
}

func NewTokenStore() *TokenStore {
	return &TokenStore{tokens: map[string]string{}}
}

func (ts *TokenStore) Write(token, email string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.tokens[token] = email
}

func (ts *TokenStore) IsValid(token string) bool {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	_, ok := ts.tokens[token]
	return ok
}

func (ts *TokenStore) Expire(token string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	delete(ts.tokens, token)
}
