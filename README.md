GW is a webserver that provides basic git functionalities via HTTP endpoints.

The most standard way to run this project is via docker (unless you're doing project development).

1. Go to the folder that you want to expose via git.
1. Copy `.env.example` and save it as `.env`
1. Copy docker-compose.yaml to your folder.
1. Update the environment variables as needed. See below on how to set a github token.
1. Run docker compose up.

**Environment Variables**

Check the sample `.env.example` file for what kinds of environment variables you can use.

**Development**

1. Clone this repo.
1. Have go mentioned in `go.mod` or higher installed.
1. In the root directory of this project, run `go install .`
1. Navigate to the folder that you want to expose through this server.
1. Run `git init` if this folder isn't a git repo already.
1. Create `.env` file in the root of this folder, mentioning the absolute path to this folder as value to the variable `GW_REPO`.
1. Optionally, set `HOST` and `PORT` environment variables as necessary.
1. Run `gw`.

## Generating Github Token

1. Go to <https://github.com/settings/personal-access-tokens/>
1. On the left sidebar, make sure "Fine-grained tokens" is selected.
1. Click Generate New Token.
1. Fill up the token name, description, fields etc.
1. Set the expiration to "No expiration" or whatever you prefer.
1. Select "Only select repository" and select the repo you want to give access to.
1. On the bottom box, click on add permissions select "Contents"
1. Give contents Read and Write access and save the token in the .env file in the `TOKEN` variable.
