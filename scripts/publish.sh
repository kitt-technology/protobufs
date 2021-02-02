#!/bin/bash

# Deals with adding files, tagging new version and pushing repo.
# 

CURR=$(git describe --tags `git rev-list --tags --max-count=1`)
NEXT=$(./scripts/semver.sh bump minor $CURR)
echo 'Building protobufs...';
make build &&
echo 'Committing...';
git add -A &&
git commit -m 'feat: update proto files' &&
git push origin main &&
git tag -a $NEXT -m "update protobufs $NEXT" &&
git push origin $NEXT