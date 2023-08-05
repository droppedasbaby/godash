#!/bin/bash

set -e
if [ "${CIRCLE_BRANCH}" = "main" ]; then
  git fetch --tags
  latestTag=$(git describe --tags `git rev-list --tags --max-count=1`)
  newTag=$(echo $latestTag | awk -F. '/v[0-9]+\.[0-9]+\.[0-9]+/{print $1"."$2"."$3+1}')
  echo "Creating a new release with tag: $newTag"
#  gh release create "${CIRCLE_TAG}" --generate-notes
else
  echo "This is not the main branch, so no release will be created."
  exit 1
fi
