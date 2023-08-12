docker build --tag kvs .
docker run --detach --publish 8080:8080 kvs