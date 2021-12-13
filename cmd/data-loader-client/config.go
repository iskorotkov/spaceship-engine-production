package main

type Config struct {
	AddrTCP    string
	AddrQUIC   string
	AddrNATS   string
	AddrGRPC   string
	ConnString string
	CertFile   string
	KeyFile    string
	RootCA     string
}
