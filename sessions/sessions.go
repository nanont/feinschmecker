package sessions

import "github.com/nanont/feinschmecker/lang"

type Session struct {
	Language lang.Language
}

type SessionMap map[int64]*Session

func Init() SessionMap {
	/* TODO: init from cache etc. */
	sessions := make(SessionMap)
	return sessions
}

func New() *Session {
	return &Session{}
}

func GetOrNew(sessions SessionMap, id int64) *Session {
	session, ok := sessions[id]
	if !ok {
		session = New()
		sessions[id] = session
	}

	return session
}
