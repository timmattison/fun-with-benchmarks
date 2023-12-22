#!/usr/bin/env bash

go test --bench=. ./golang-printing | uniq
