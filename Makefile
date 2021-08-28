CC 		:= GO111MODULE=on CGO_ENABLED=0 go
CFLAGS 		:= build -o
SHELL 		:= /bin/bash
MODULE 		:= github.com/ok-john/cryptopals
GO_SRC		:= https://raw.githubusercontent.com/ok-john/tmpl-go/main/install-go
TAG_SRC		:= https://raw.githubusercontent.com/ok-john/tag/main/tag

ifneq ($(MAKECMDGOALS),)
FIRST_GOAL := $(word 1, $(MAKECMDGOALS))
LAST_GOAL := $(word $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
else
FIRST_GOAL := all
LAST_GOAL := all
endif

$(FIRST_GOAL) :: link 

mod-install ::
				$(CC) install ./... 

tidy :: mod-install
				$(CC) mod tidy -compat=1.17
				
format :: tidy
				$(CC)fmt -w -s *.go

test ::	 format
				$(CC) test -v ./...

compile :: test
				$(CC) $(CFLAGS) $(MODULE) && chmod 755 $(MODULE)

link :: compile
				$(shell ldd $(MODULE))

headers :: link
				$(shell readelf -h $(MODULE) > $(MODULE).headers)

copy-up :: 
				cp $(MODULE) .

run ::  copy-up
				./$(MODULE)

clean :: 
				rm -rf github.com

$(LAST_GOAL) :: run

install-scripts :: 
				cat <(curl -sS $(TAG_SRC)) > tag && chmod 755 tag
				cat <(curl -sS $(GO_SRC)) > install-go && chmod 755 install-go
				
trace :: 
				trace-cmd record -p function_graph -F ./$(MODULE)
				trace-cmd report | sed 's/.* | //g' > $(MODULE).ttree
				trace-cmd record -p function_graph -e syscalls -F ./$(MODULE)
				trace-cmd report | sed 's/.* | //g' > $(MODULE)-syscalls.ttree
				trace-cmd record -p function_graph -g __x64_sys_read ./$(MODULE)
				trace-cmd report | sed 's/.* | //g' > $(MODULE)-sysreads.ttree
