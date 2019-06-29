# Fleet

*Work in Progress*

## Installation
```
$ go get github.com/seadreamcode/fleet
```

## Getting Started

Fleet will look for the `.fleet.yml` file in the current
working directory.
```yaml
# .fleet.yml example
tasks:
    build:
        - npm run build
        - npm run clean
    test:
        - npm run test
    deploy:
        - npm run deploy
hooks:
    github:
        repo: myorg/repo
        steps:
            - build
            - test
            - deploy    
```

This example script defines 3 tasks that are then sequenced by a webhook from Github.

## Using the CLI

Start the fleet daemon
```
$ fleet up
```

Stop the fleet daemon
```
$ fleet down
```

Observe the health of the local node
```
$ fleet health
```

Fetch the container logs
```
$ fleet logs
```