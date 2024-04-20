package token

import "time"

type Maker interface {
	//Creates token for specific token and duration
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}