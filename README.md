# mirbase
[<img src="https://img.shields.io/docker/image-size/thavlik/mirbase/latest">](https://hub.docker.com/r/thavlik/mirbase)
[<img src="https://img.shields.io/badge/maintenance%20status-actively%20developed-brightgreen">](https://github.com/thavlik/mirbase)
[<img src="https://img.shields.io/badge/Language-go-01add8.svg">](https://go.dev/)
[<img src="https://img.shields.io/badge/License-MIT-lightblue.svg">](./LICENSE)

This repository contains a Dockerfile for building a [sqlite](https://www.sqlite.org/) database from the [miRBase data](https://www.mirbase.org/download/). 

## Building
Run the following command:
```bash
docker build -t thavlik/mirbase:latest .
```

## Schema
See [tables.sql](pkg/store/sql_store/tables.sql) for how the tables are created. Note that the schema used for sqlite differs from the official miRBase release.

### Searching
[tables.sql](pkg/store/sql_store/tables.sql) creates an [fts5](https://www.sqlite.org/fts5.html) virtual table to fuzzy search for rows in the `mirna` table. The trigram tokenizer provides a fuzzy searching behavior. 

## Usage
The prebuilt image is intended to be used as a base image, where you can find the database file at `/mirbase.sqlite`. For example:

```Dockerfile
# Create a reference to the database image as a build stage.
FROM thavlik/mirbase:latest AS db

# Create your image here.
FROM debian:latest
COPY --from=db /mirbase.sqlite /opt/mirbase.sqlite
#
# Now your docker image contains the prebuilt
# sqlite database file at /opt/mirbase.sqlite
#
```

Note: the only time you must build this image is if you want to change how data is handled. For example, you want to insert into a PostgreSQL server instead of writing to a sqlite file. If sqlite is sufficient for you, you are encouraged to use the [prebuilt image](https://hub.docker.com/r/thavlik/mirbase).

## License
The code in this repository for building the database is [licensed under MIT](./LICENSE).