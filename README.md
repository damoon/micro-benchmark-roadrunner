# roadrunner, apache, golang "hello world" benchmark

# redo

```
docker-composer up --build
```

```
ab -c 1000 -n 100000 http://127.0.0.1:8080/
sleep 10
ab -c 1000 -n 100000 http://127.0.0.1:8080/ > benchmark/hello-world-roadrunner.txt
ab -c 1000 -n 100000 http://127.0.0.1:8081/
sleep 10
ab -c 1000 -n 100000 http://127.0.0.1:8081/ > benchmark/hello-world-apache-php.txt
ab -c 1000 -n 100000 http://127.0.0.1:8082/
sleep 10
ab -c 1000 -n 100000 http://127.0.0.1:8082/ > benchmark/hello-world-golang.txt
```
