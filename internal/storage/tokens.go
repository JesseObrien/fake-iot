package storage

import (
	"sync"
)

type UserIdentity struct {
	email     string
	accountId string
}

type TokenStore struct {
	tokens map[string]UserIdentity
	mu     sync.Mutex
}

func NewTokenStore() *TokenStore {
	return &TokenStore{tokens: map[string]UserIdentity{}}
}

func (ts *TokenStore) Write(token, email, accountId string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.tokens[token] = UserIdentity{email, accountId}
}

func (ts *TokenStore) IsValid(token string) bool {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	_, ok := ts.tokens[token]
	return ok
}

func (ts *TokenStore) IsValidAccountToken(token string, accountId string) bool {

	if !ts.IsValid(token) {
		return false
	}

	ts.mu.Lock()
	defer ts.mu.Unlock()

	identity := ts.tokens[token]

	return identity.accountId == accountId
}

func (ts *TokenStore) Expire(token string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	delete(ts.tokens, token)
}
