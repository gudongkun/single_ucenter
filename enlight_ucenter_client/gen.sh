#!/bin/bash

for X in $(find . -name "*.proto" | sed "s|^\./||"); do
	protoc -I$(pwd) --go_out=paths=source_relative:.  --micro_out=paths=source_relative:. $X
done