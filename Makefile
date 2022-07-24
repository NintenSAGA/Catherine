
GO_DIR = go_version/
C_DIR  = c_version/

run_go: $(MAKEFILE_GO)
	cd $(GO_DIR) && make run

run_c: $(MAKEFILE_C)
	cd $(C_DIR) && make run