package model

import "io"

type Context struct {
	Code int
	Url  string
	Body io.ReadCloser
}
