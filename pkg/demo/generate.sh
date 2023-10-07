#!/bin/bash

# To compile all protobuf files in this repository, run
# "make protobuf" at the top-level.

protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  *.proto
