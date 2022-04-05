Go Lang Goose
- [Goose Documentation](https://github.com/pressly/goose)

install goose CLI
```sh
$ curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh |    GOOSE_INSTALL=$HOME/.goose sh

$ alias goose=”$HOME/.goose/bin/goose”
```

If you want to make alias permanent, do the following steps:
```sh
$ sudo nano ~/.bash_aliases
```
Paste the following text

```
#My custom aliases
alias goose="$HOME/.goose/bin/goose"
```

Reload the bash
```sh
source ~/.bashrc
```

Create a new migration
```sh
goose --dir database/goose/migrations create init_schema sql
```

See migrations status
```sh
goose --dir database/goose/migrations postgres "postgresql://postgres_username:postgres_password@localhost:5432/postgres_database?sslmode=disable" status
```

Run up migrations
```sh
goose --dir database/goose/migrations postgres "postgresql://postgres_username:postgres_password@localhost:5432/postgres_database?sslmode=disable" up
```

Run down migrations
```sh
goose --dir database/goose/migrations postgres "postgresql://postgres_username:postgres_password@localhost:5432/postgres_database?sslmode=disable" down
```

