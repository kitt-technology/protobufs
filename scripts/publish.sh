#!/bin/bash


CURR=$(git describe --tags `git rev-list --tags --max-count=1`)


NEXT=$(./scripts/semver.sh bump minor $CURR)

echo 'Building protobufs...';
make build &&
echo 'Committing...';
git add -A &&
git commit -m 'feat: update proto files' &&

git tag -a $NEXT -m "update protobufs $NEXT"
git push origin $NEXT