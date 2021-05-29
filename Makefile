DEST_DIR := "build"
GO := "go"
CMD := "slucg"

.PHONY: build clean

build:
	mkdir -p $(DEST_DIR)
	$(GO) build -o $(DEST_DIR) ./...

clean:
	rm -rf $(DEST_DIR)
