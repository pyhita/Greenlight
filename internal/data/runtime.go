package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeData = errors.New("invalid runtime data")

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quoteValue := strconv.Quote(jsonValue)

	return []byte(quoteValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	res, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeData
	}

	parts := strings.Split(res, " ")
	if len(parts) < 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeData
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeData
	}

	*r = Runtime(i)
	return nil
}
