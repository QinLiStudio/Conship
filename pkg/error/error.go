//Author: lxk20021217
//Date: 2022-06-21 21:58:59
//LastEditTime: 2022-06-24 14:49:20
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\pkg\error\error.go
//是谁总是天亮了才睡

package error

import "github.com/pkg/errors"

var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	Unwrap       = errors.Unwrap
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	Cause        = errors.Cause
	Errorf       = errors.Errorf
	Is           = errors.Is
	As           = errors.As
)

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}
