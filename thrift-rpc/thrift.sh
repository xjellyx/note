#!/usr/bin/env bash
thrift -r --gen go:package_prefix=github.com/olongfen/node/thrift-rpc/gen-go/ demo.thrift
