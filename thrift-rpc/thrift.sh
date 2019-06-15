#!/usr/bin/env bash
thrift -r --gen go:package_prefix=github.com/srlemon/node/thrift-rpc/gen-go/ example.thrift
