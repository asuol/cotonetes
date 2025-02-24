# Cotonetes - webpage note organizer

WIP

## Dockerized golang development environment

Build the golang based docker image with `make build`

Then create the following environment variables, each pointing to the path of an empty directory:

`COTONETES_GOCACHE` - Persist the project build dependencies (specially the sqlite CGO build)
`COTONETES_GOMODCACHE` - Persist the project dependencies (to avoid downloading on every container run)

Then run either `make sh` to open a shell ou run `make run`

You may also just use the form

`COTONETES_GOCACHE=/your/path/here COTONETES_GOMODCACHE=/your/path/here make run`
