reinstall_cmd_misa:
	go install ./cmd/misa && misa ls

test: modvendor
	go test ./... -v --cover

modvendor:
	GOPROXY=https://goproxy.cn go mod vendor

default: reinstall_cmd_misa
