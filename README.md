# mirbase
[<img src="https://img.shields.io/badge/maintenance%20status-actively%20developed-brightgreen">](https://github.com/thavlik/mirbase)
[<img src="https://img.shields.io/badge/Language-go-01add8.svg">](https://go.dev/)
[<img src="https://img.shields.io/badge/License-MIT-lightblue.svg">](./LICENSE)

This repository contains a Dockerfile for building a sqlite database from the [miRBase data](https://www.mirbase.org/download/).

## Building
Run the following command:
```bash
docker build -t thavlik/mirbase:latest .
```

## Usage
This image is intended to be used as a base image. For example:

```Dockerfile
# Create a reference to the database image as a build stage.
FROM thavlik/mirbase:latest AS db

# Create your image here.
FROM debian:latest
COPY --from=db /mirbase.sqlite /opt/mirbase.sqlite
#
# Now your docker image contains the prebuilt sqlite database file.
#
```

## License
The code in this repository for building the database is [licensed under MIT](./LICENSE).