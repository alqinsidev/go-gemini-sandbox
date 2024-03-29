package error

import "errors"

var (
	ErrBadPayloadRequest         = errors.New("bad body payload")
	ErrDB                        = errors.New("database error")
	ErrNotFound                  = errors.New("entity not found")
	ErrNeedToPopulateInformation = errors.New("you need to add information first")
)
