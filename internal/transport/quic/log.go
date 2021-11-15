package quic

import (
	logpkg "log"
	"os"
)

var log = logpkg.New(os.Stderr, "[quic] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)
