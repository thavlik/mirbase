# mirbase
This repository contains a Dockerfile for building a sqlite database from the [miRBase data](https://www.mirbase.org/download/).

## Building
Run the following command:
```bash
docker build -t thavlik/mirbase:latest .
```

## Using
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