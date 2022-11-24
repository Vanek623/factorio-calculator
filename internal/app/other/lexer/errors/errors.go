package errors

import "github.com/pkg/errors"

var ErrFieldParse = errors.New("cannot parse field")
var ErrIncorrectHandler = errors.New("this handler cannot parse this field")
