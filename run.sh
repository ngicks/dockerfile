#!/bin/bash

deno run --allow-read=.env,.env.defaults --allow-run --allow-env script/docker_run.ts $@