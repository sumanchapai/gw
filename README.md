GW is a webserver that provides basic git functionalities via HTTP endpoints.

**Motivation:**

For small, single user web apps, I use text-based databases. For example, JSON,
or YAML files. Similarly, the accounting software that I use, beancount, is
also text-based. Some people ask, why YAML file as a database? Where
performance isn't an issue, particularly in single-user apps, having text files
as database means that your database is amenable to version control (simply
using git) and has high fault tolerance by simply pushing your text file to
github. However, the users of these web apps aren't all computer programmers.
For example, one of the web app that I recently made is for our receptionist to
create/update/delete monthly analytics information for our guesthouse. For
example, our current review on Google, number of new reviews, etc. etc. We need
to note this information every month to track our performance and draw some
charts from them to see progress. I wrote a frontend in react router v7 and
have JSON file as a database served by `json-server`. I use this app, `gw` so
that the receptionist can simply make commits by clicking some buttons in some
website. The buttons under the hood run commands like `git commit` and `git
push` but the receptionist doesn't need to know that, or rather wouldn't want
to know that. They can simply know that those buttons "SAVES" their data to server
and it can be recovered even if our computer went down. Having text based database,
coupled with git, is incredibly useful also when your database is saving data
that you would not want tampered, intentionally or otherwise. For example, in an
accounting software, you wouldn't want your accountant to change the previous month
salary of some employee. In plain-text database, that's easily caught by just using
`git diff`. You can always review all the commits of the your accountant and catch any
mistakes easily.

![Web view of `gw` app](images/image1.png)

**Usage:**

The most standard way to run this project is via docker (unless you're doing project development).

1. Go to the folder that you want to run git commands from Web browser.
1. Copy `.env.example` from this project repo and save it as `.env` into your folder.
1. Copy `docker-compose.yaml` from this project repo and save it to your folder.
1. Update the environment variables as needed. See below on how to set a github token.
1. Run docker compose up.
1. Visit [localhost:8071](http://localhost:8071)

**Environment Variables:**

Check the sample `.env.example` file for what kinds of environment variables you can use.

**Development:**

1. Clone this repo.
1. Have go mentioned in `go.mod` or higher installed.
1. In the root directory of this project, run `go install .`
1. Navigate to the folder that you want to expose through this server.
1. Run `git init` if this folder isn't a git repo already.
1. Create `.env` file in the root of this folder, mentioning the absolute path to this folder as value to the variable `GW_REPO`.
1. Optionally, set `HOST` and `PORT` environment variables as necessary.
1. Run `gw`.

**Generating Github Token:**

1. Go to <https://github.com/settings/personal-access-tokens/>
1. On the left sidebar, make sure "Fine-grained tokens" is selected.
1. Click Generate New Token.
1. Fill up the token name, description, fields etc.
1. Set the expiration to "No expiration" or whatever you prefer.
1. Select "Only select repository" and select the repo you want to give access to.
1. On the bottom box, click on add permissions select "Contents"
1. Give contents Read and Write access and save the token in the .env file in the `TOKEN` variable.
