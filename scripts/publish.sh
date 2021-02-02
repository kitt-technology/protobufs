#!/bin/bash

echo 'Building protobufs...';
make build &&

echo 'Committing...';
git add -A &&
git commit -m 'feat: update proto files' &&

