# How to run?

`$ docker run --publish 3000:8080 --name goweb golang-docker` or
`$ docker run --publish 3000:8080 --name goweb --rm golang-docker` // It will remove the container when it exists.

# Start and stop the container
`$ docker stop goweb`
`$ docker run goweb`