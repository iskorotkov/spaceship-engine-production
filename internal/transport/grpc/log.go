package grpc

import (
	logpkg "log"
	"os"
)

var log = logpkg.New(os.Stderr, "[grpc] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)
