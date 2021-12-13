protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/report-printer/report-printer.proto

	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/data-loader/data-loader.proto

increase_udp_buffer:
	sysctl -w net.core.rmem_max=2500000
