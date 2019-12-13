reinstall_cmd_misa:
	go install ./cmd/misa && misa ls

test: modvendor
	go test ./... -v --cover

modvendor:
	GO111MODULE=on GOPROXY=https://goproxy.cn go mod vendor

gitpush: test
	git push

default: reinstall_cmd_misa
