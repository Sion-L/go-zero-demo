#!/bin/bash
function start_demo() {
    go run demo/demo.go -f demo/etc/demo.yaml
}

function start_transform() {
    go run transform/transform.go -f transform/etc/transform.yaml
}

function start_gateway() {
    go run gateway/gateway.go -f gateway/etc/gateway.yaml
}

(start_demo) &
(start_transform) &
(start_gateway)