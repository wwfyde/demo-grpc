#!/bin/bash
# .proto file directory

# get PROJECT_DIR
PROJECT_DIR="$( cd "$( dirname $(dirname $0) )" && pwd )"
cd "$PROJECT_DIR/examples/helloworld"
echo $(pwd)

