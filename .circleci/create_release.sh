Copy code
#!/bin/bash

if [ "${CIRCLE_BRANCH}" = "master" ]; then
  git fetch --tags
  gh auth login --with-token < ${GH_TOKEN}
  gh release create ${CIRCLE_TAG} --title "${CIRCLE_TAG}" --notes "Release ${CIRCLE_TAG}"
fi
