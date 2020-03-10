#!/usr/bin/env bash
# Check argument count
if [ $# -ne 1 ]; then
  echo "ARGS: $#" 1>&2
  echo "Error: Require argument=1" 1>&2
  exit 1
fi

cd $(dirname $(cd $(dirname $0); pwd))

ECR_REPOSITORY_URL=$1

# AWS Login
$(aws ecr get-login --no-include-email)

# Build
ECR_REPOSITORY_URL=$ECR_REPOSITORY_URL docker-compose -f docker-compose.prod.yml build
docker push "$ECR_REPOSITORY_URL":latest

exit 0
