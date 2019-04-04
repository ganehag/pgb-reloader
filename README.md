# PgBouncer Reloader

Reloads PgBouncer whenever configuration changes.

For use with my other PgBouncer project: https://github.com/ganehag/docker-pgbouncer

# Configuration

The software takes two Environment variables.

- PGBOUNCER_CONFIG
- PGBOUNCER_URI


## PGBOUNCER_CONFIG

A semicolor ';' separated list of files.

Example: `"/etc/pgbouncer/databases.txt;/etc/pgbouncer/pgbouncer.ini`


## PGBOUNCER_URI

Connection string to access PGBouncers admin interface.

Note: Username should match `admin_users` in the main PgBouncer config.

Example: postgresql://postgres:secret@127.0.0.1:6432/pgbouncer?sslmode=disable ./pgb-reload 


# How it works

A simple Golang application watches one or several files for changes.
When a change is detected, it connects to PgBouncers admin interface
and issues a `RELOAD` command.

```
export PGBOUNCER_CONFIG="/etc/pgbouncer/databases.txt"
export PGBOUNCER_CONFIG="postgresql://postgres:secret@127.0.0.1:6432/pgbouncer?sslmode=disable"
./pgb-reload 
2019/04/04 18:13:16 INFO Watching '/etc/pgbouncer/databases.txt'
```

Then when `databases.txt` changes:

```
2019/04/04 18:14:59 INFO change event
```

And in the PgBouncer log:

```
2019-04-04 18:13:16.475 6 LOG RELOAD command issued
```


# Kubernetes

pgb-reloader can easily be added to a Kubernetes pod as a sidecard.

```yaml
write example yaml
```
