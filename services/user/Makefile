PROTOBUF_DIR=../../protobuf
OUT_DIR=pkg/proto/userpb

proto_clean:
	@echo "Cleaning up generated Go gRPC code..."
	@if [ -d $(OUT_DIR) ]; then \
		find $(OUT_DIR) -mindepth 1 -delete; \
	else \
		echo "$(OUT_DIR) does not exist, nothing to clean."; \
	fi
	@echo "Clean up completed."

proto: proto_clean
	@echo "Creating $(OUT_DIR) directory if it does not exist..."
	@mkdir -p $(OUT_DIR)
	@echo "Compiling gRPC code for all .proto files in $(PROTOBUF_DIR)/userpb"
	@find $(PROTOBUF_DIR)/userpb -name '*.proto' | xargs -I {} protoc -I=$(PROTOBUF_DIR) --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
                                                            --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative {} \

	@echo "gRPC compilation completed."