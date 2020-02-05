#!/usr/bin/env bash
thrift -r --gen go:package_prefix=github.com/olongfen/node/strings-learn/gen-go/ demo.thrift
