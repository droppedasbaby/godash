#!/bin/bash

set -e
if [ "${CIRCLE_BRANCH}" = "main" ]; then
  git fetch --tags
  echo "${GITHUB_TOKEN}" | gh auth login --with-token
  gh release create ${CIRCLE_TAG} --title "${CIRCLE_TAG}" --notes "Release ${CIRCLE_TAG}"
else
  echo "This is not the main branch, so no release will be created."
  exit 1
fi
