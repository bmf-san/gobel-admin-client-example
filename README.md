# gobel-admin-client-example
[![GitHub license](https://img.shields.io/github/license/bmf-san/gobel-admin-client-example)](https://github.com/bmf-san/gobel-admin-client-example/blob/master/LICENSE)
[![CircleCI](https://circleci.com/gh/bmf-san/gobel-admin-client-example.svg?style=svg)](https://circleci.com/gh/bmf-san/gobel-admin-client-example)

Gobel is a headless cms built with golang. 

This is an example for admin client application of gobel.

# gobel
- [bmf-tech-client](https://github.com/bmf-san/bmf-tech-client)
- [gobel-api](https://github.com/bmf-san/gobel-api)
- [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)
- [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
- [gobel-example](https://github.com/bmf-san/gobel-example)
- [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)
- [migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel)

# Dockerhub
[bmfsan/gobel-admin-client-example](https://hub.docker.com/r/bmfsan/gobel-admin-client-example)

# Get started
Before you start, you need to clone [gobel-api](https://github.com/bmf-san/gobel-api).

## Docker
### Create an env file
```
cp .env_example .env
```

### Edit a host file
```
127.0.0.1 gobel-admin-client-example.local
```

### Build and run containers
```
make docker-compose-up

or

make docker-compose-up-d
```

Then go to `gobel-admin-client-example.local:82`

## npm
### Create an .env.local
```
cp .env /app/.env.local
```

### Starts a dev server with hot-reload
```
cd app/
npm install
npm run serve
```

## npm scripts
|   Scripts   |             Description              |
| ----------- | ------------------------------------ |
| serve       | Starts a dev server with hot-reload. |
| local-build | Run build for local.                 |
| prod-build  | Run build for production.            |
| test:unit   | Run unit tests.                      |
| lint        | Run lint with --fix option.          |

# Contributing
We welcome your issue or pull request from everyone.
Please make sure to read the [CONTRIBUTING.md](https://github.com/bmf-san/gobel-admin-client-example/.github/CONTRIBUTING.md).

# License
This project is licensed under the terms of the MIT license.

# Author
bmf - Software engineer.

- [github - bmf-san/bmf-san](https://github.com/bmf-san/bmf-san)
- [twitter - @bmf-san](https://twitter.com/bmf_san)
- [blog - bmf-tech](http://bmf-tech.com/)
