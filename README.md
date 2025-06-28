## gator
gator is an RSS feed aggregator.

You can:
- Add RSS feeds from across the internet to be collected.
- Store the collection posts in a PostgreSQL database.
- Follow RSS feeds that other users have added.
- View summaries and meta data about posts in the terminal.

## Dependencies:
- go v1.24
- PostgreSQL v14.18

## Setup

To build gator 
```bash

$ go build

```

You also need to create a database of the name of 'gator' in you 
postgresql

You need a `.gatorconfig.json` file in home directory. This is how
gator manages user sessions and database connection,
```bash

touch  ~/.gatorconfig.json

```

Your config file needs to have two fields 'db_url' which is connection
URL to your postgres database and 'currect_user_name' which is current 
logged in user.
Your config file should look something like this where you add your 
database link
```json

{
    "db_url":"postgres://username:password@host:port/database?sslmode=disable",
    "current_user_name":""
}

```


## Commands

- For user registration
```bash

$ gator register naveed

```

- For user login
```bash

$ gator login naveed

```


- To list all the users in the database
```bash

$ gator users

```

- To add RSS feed
```bash

$ gator addfeed [feed-name] [feed-url]

```

- To aggregate feeds.
(This is running command and will fetch the feed data in a given interval, you can still
use gator in a different terminal window)

**Caution**
Do not DOS the servers you are fetching feeds from, make sure the [time-request-request]
is reasonable.

```bash

$ gator agg [time-between-request]

```

- List feeds added by users
```bash

$ gator feeds

```

- Follow feeds added by users
```bash

$ gator follow [url]

```

- List feeds followed by particular user
```bash

$ gator following

```

- Unfollow feeds added by users
```bash

$ gator unfollow [feed-name]

```

- Browse following feeds
```bash

$ gator browse [limit] # limit by default is 2

```
