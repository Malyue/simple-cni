# Define the list of plugins to build
PLUGINS := bridge ipam

# Output directory for binaries
BIN_DIR := bin

# Target platform (Linux/amd64)
GOOS := linux
GOARCH := amd64

# Default target
all: $(PLUGINS)

# Rule to build each plugin
$(PLUGINS):
	@mkdir -p $(BIN_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$@ ./plugins/$@

# Clean up
clean:
	@rm -rf $(BIN_DIR)

# Run all unit tests
test:
	go test -v ./...

cp:
	make bridge
	mv bin/bridge bin/bridge_per
	minikube cp bin/bridge_per /opt/cni/bin/

.PHONY: all $(PLUGINS) clean test