package tcp

import (
	logpkg "log"
	"os"
)

var log = logpkg.New(os.Stderr, "[tcp] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)
