SOURCE="${BASH_SOURCE[0]}"
SCRIPT_DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
SOURCE_DIR="$(readlink -f ${SCRIPT_DIR}/..)"

cd $SOURCE_DIR

mkdir -p data/mysqldb
mkdir -p data/mysqldb_rep1
mkdir -p data/kafka
mkdir -p data/zookeeper

docker build -f Dockerfile.service . -t vinai-video-analytics-hls-streaming-service:1.1