package crossconnect

//go:generate protoc -I . crossconnect.proto --go_out=plugins=grpc:. --proto_path=$GOPATH/src
