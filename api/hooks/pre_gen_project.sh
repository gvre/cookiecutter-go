#!/usr/bin/env bash

set -eu -o pipefail

check_requirements() {
  local reqs=("go" "git" "pre-commit")
  local res=0

  for req in "${reqs[@]}"; do
    if ! command -v "${req}" &> /dev/null; then
      echo
      echo "### ERROR ###  Requirement ${req} does not exist."
      echo "Please see README.md for the list of requirements and links for installation instructions."
      echo
      res=1
    fi
  done

  return ${res}
}

check_requirements