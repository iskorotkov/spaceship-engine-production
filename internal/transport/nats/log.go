package nats

import (
	logpkg "log"
	"os"
)

var log = logpkg.New(os.Stderr, "[nats] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)
