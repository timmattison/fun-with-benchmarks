#!/usr/bin/env bash

go test --bench=. | uniq
