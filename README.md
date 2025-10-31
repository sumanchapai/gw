GW is a webserver that provides basic git functionalities via HTTP endpoints.

**Environment Variables**

Can be defined in the `.env` file.

- `GW_REPO`: absolute path to the git repo that you want to expose through this server.
- `PORT`: port that this server should run on.
- `HOST`: hostname that this server should run on.

**Development**

1. Clone this repo.
1. Have go mentioned in `go.mod` or higher installed.
1. In the root directory of this project, run `go install .`
1. Navigate to the folder that you want to expose through this server.
1. Run `git init` if this folder isn't a git repo already.
1. Create `.env` file in the root of this folder, mentioning the absolute path to this folder as value to the variable `GW_REPO`.
1. Optionally, set `HOST` and `PORT` environment variables as necessary.
1. Run `gw`.
1. Note: A bug exists currently where in if you run `gw` from a different location, some git actions will assume that location to be the location to git repo ignoring `GW_REPO` environment variable.
