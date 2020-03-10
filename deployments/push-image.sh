#!/usr/bin/env bash
# Check argument count
if [ $# -ne 1 ]; then
  echo "ARGS: $#" 1>&2
  echo "Error: Require argument=1" 1>&2
  exit 1
fi

cd $(dirname $(cd $(dirname $0); pwd))

ARG_REPOS_URL=$1

# AWS Login
$(aws ecr get-login --no-include-email)

# Build
ARG_REPOS_URL=$ARG_REPOS_URL docker-compose -f docker-compose.prod.yml build
docker push "$ARG_REPOS_URL":latest

exit 0
