# Build a runable binary for the Linux container

`$ CGO_ENABLED=0 GOOS=linux go build -o main .`

# Build the docker image

`$ docker build -t golang-scratch .`

# Run it

`$ docker run -it golang-scratch` or
`$ docker run -it --name scratch --rm golang-scratch`
