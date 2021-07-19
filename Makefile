DEST_DIR := "build"
GO := "go"
CMD := "ackslack"

.PHONY: build clean

build:
	mkdir -p $(DEST_DIR)
	$(GO) build -o $(DEST_DIR) ./...

clean:
	rm -rf $(DEST_DIR)
