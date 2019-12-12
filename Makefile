reinstall_cmd_misa:
	go install ./cmd/misa && misa ls

default: reinstall_cmd_misa