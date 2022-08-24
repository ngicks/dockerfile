package main

//go:generate go run ./fill_base.go -o ../deno_node_dev/Dockerfile -option ./config/deno_node.json
//go:generate go run ./fill_base.go -o ../dood_local/Dockerfile -option ./config/dood_local.json
//go:generate go run ./fill_base.go -o ../go_dev/Dockerfile -option ./config/go.json
//go:generate go run ./fill_base.go -o ../rust_dev/Dockerfile -option ./config/rust.json
