year=$1
day=$2
puzzle=$3

echo "Running for $year / $day-$puzzle"
DOCKER_BUILDKIT=0 docker build . -t advent-of-code

echo "Starting container.... for  $year / $day-$puzzle"
docker run advent-of-code  ./start.bash $year $day $puzzle