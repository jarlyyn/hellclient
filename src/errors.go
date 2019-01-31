package main

import "errors"

var errFuncWhenRunFuncNotRewrited = func() {
	panic(errors.New("no application run funciton.You should rewrite default run function"))
}
