#!/usr/bin/env bash
thrift -r --gen go:package_prefix=github.com/LnFen/node/thrift-rpc/gen-go/ demo.thrift