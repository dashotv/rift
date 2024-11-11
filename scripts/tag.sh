#!/usr/bin/env bash

set -e

tag="$1"
if [[ -z $tag ]]; then
  echo "Usage: $0 <tag>"
  exit 1
fi

if [[ "$tag" == $(git tag -l "$tag") ]]; then
  echo "Tag $tag already exists"
  exit 1
fi

if [[ -n "$(git status -s)" ]]; then
  echo "There are pending changes, commit first"
  exit 1
fi

cat >internal/scraper/version.go <<EOF
package scraper

// VERSION is the version of the application
// automatically managed by pre-commit githook
const VERSION = "$tag"
EOF
git add internal/scraper/version.go
git commit -m "$tag"
git tag -a "$tag" -m "$tag"
