#!/usr/bin/env bash
thrift -r --gen go:package_prefix=github.com/srlemon/node/strings-learn/gen-go/ demo.thrift
