package lib

import "github.com/rs/xid"

func SecureRandom() (string, error) {
	return xid.New().String(), nil
}
