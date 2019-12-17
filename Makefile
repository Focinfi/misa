reinstall_cmd_misa:
	go install ./cmd/misa && misa ls

test: modvendor
	go test ./... -v --cover

modvendor:
	GO111MODULE=on GOPROXY=https://goproxy.cn go mod vendor

gitpush: test
	git push

testHttp:
	misa run get-first-pipeline-id -d '{"method": "GET", "url": "https://raw.githubusercontent.com/Focinfi/misa/master/configs/conf.example.json"}'

default: reinstall_cmd_misa
