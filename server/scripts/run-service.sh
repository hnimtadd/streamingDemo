SOURCE="${BASH_SOURCE[0]}"
SCRIPT_DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
SOURCE_DIR="$(readlink -f ${SCRIPT_DIR}/..)"

cd $SOURCE_DIR

docker-compose -f docker-compose.yml up -d
