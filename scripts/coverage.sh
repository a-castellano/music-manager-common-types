#!/bin/bash

mkdir -p cover
echo 'mode: count' > cover/coverage.report

PKG_LIST=$(go list ./... | grep -v /vendor/)
for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "cover/${package##*/}.cov" "$package" ;
done
tail -q -n +2 cover/*.cov >> cover/coverage.report
go tool cover -func=cover/coverage.report
go tool cover -html=cover/coverage.report -o coverage.html
