Go Lang Migrate
- [Migrate Documentation](https://github.com/golang-migrate/migrate)

install migrate CLI
```sh
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

create a new migration
```sh
$ migrate create -ext sql -dir database/migrations -seq [migration_name]
```

run up migrations

```sh
$ migrate -path database/migrations -database "postgresql://postgres_username:postgres_password@localhost:5432/postgres_database?sslmode=disable" -verbose up
```

run down all migration

```sh
$ migrate -path database/migrations -database "postgresql://postgres_username:postgres_password@localhost:5432/postgres_database?sslmode=disable" -verbose down -all 
```

run down specific migration, it is important to remember the migration_version argument is the number of how many migration you want down of the latest versions

```sh
$ migrate -path database/migrations -database "postgresql://postgres_username:postgres_password@localhost:5432/postgres_database?sslmode=disable" -verbose down [migration_version]
```