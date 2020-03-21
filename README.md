# opype

## Tools

This project uses golint so please make sure you install it and have it in your path

`
go get -u golang.org/x/lint/golint
`

## Start and stop the developer infrastructure

We need a jenkins instance for now that will just run a jenkins instance. 
We can than create as many projects as we want and get logs from them.

`
make start-dev-infra
`
will start the docker-compose file

`
make stop-dev-infra
`
will stop it