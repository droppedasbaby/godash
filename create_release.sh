#!/bin/bash

set -e
if [ "${CIRCLE_BRANCH}" = "main" ]; then
  git fetch --tags
  latestTag=$(git describe --tags `git rev-list --tags --max-count=1`)
  case $INCREMENT_LEVEL in
    "major")
      newTag=$(echo $latestTag | awk -F. '/v[0-9]+\.[0-9]+\.[0-9]+/{print $1+1"."$2"."$3}')
      ;;
    "minor")
      newTag=$(echo $latestTag | awk -F. '/v[0-9]+\.[0-9]+\.[0-9]+/{print $1"."$2+1"."$3}')
      ;;
    "patch")
      newTag=$(echo $latestTag | awk -F. '/v[0-9]+\.[0-9]+\.[0-9]+/{print $1"."$2"."$3+1}')
      ;;
    *)
      echo "Invalid increment level. Must be one of 'major', 'minor', 'patch'"
      exit 1
  esac
  echo "Creating a new release with tag: $newTag"
#  gh release create "${CIRCLE_TAG}" --generate-notes
else
  echo "This is not the main branch, so no release will be created."
  exit 1
fi
