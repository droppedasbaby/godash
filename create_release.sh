#!/bin/bash

set -e
if [ "${CIRCLE_BRANCH}" = "main" ]; then
  gh release create "${CIRCLE_TAG}" --generate-notes
else
  echo "This is not the main branch, so no release will be created."
  exit 1
fi
