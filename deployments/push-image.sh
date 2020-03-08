# Check argument count
if [ $# -ne 2 ]; then
  echo "ARGS: $#" 1>&2
  echo "Error: Require argument=2" 1>&2
  exit 1
fi

cd $(dirname $(cd $(dirname $0); pwd))

ARG_PROFILE=$1
ARG_REPOS_URL=$2

# AWS Login
$(aws ecr get-login --no-include-email --profile ${ARG_PROFILE})

# Build
docker-compose -f docekr-compose.prod.yml build
docker push "$ARG_REPOS_URL":latest

exit 0