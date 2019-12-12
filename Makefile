reinstall_cmd_misa:
	go install ./cmd/misa && misa ls

test: 
	go test ./... -v --cover

default: reinstall_cmd_misa
