package fmts

import (
	gconcat "github.com/jeffotoni/gconcat"
)

//Concat contaquena
func Concat(strs ...interface{}) string {
	return gconcat.Build(strs...)
}
