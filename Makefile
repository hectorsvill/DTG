# List of target platforms
PLATFORMS := windows/amd64 linux/amd64 darwin/amd64

# Output file name
OUTPUT_FILE := dir_tree.zip

# Temporary directory for binary files
TEMP_DIR := .bin

# Build the code for all target platforms and create the output file
build: $(PLATFORMS)
	zip -j $(OUTPUT_FILE) $(TEMP_DIR)/*

# Rule to build the code for a specific platform
$(PLATFORMS):
	$(eval GOOS=$(word 1,$(subst /, ,$@)))
	$(eval GOARCH=$(word 2,$(subst /, ,$@)))
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(TEMP_DIR)/dir_tree-$(GOOS)-$(GOARCH)

# Create the temporary directory
$(TEMP_DIR):
	mkdir -p $(TEMP_DIR)

# Clean up the build artifacts
clean:
	rm -rf $(TEMP_DIR) $(OUTPUT_FILE)

.PHONY: build $(PLATFORMS) clean