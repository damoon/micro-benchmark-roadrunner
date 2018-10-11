# roadrunner, apache, golang "hello world" benchmark

# redo

```
docker-composer up --build
```

```
ab -c 1 -t 30 http://127.0.0.1:8080/
sleep 10
ab -c 1 -t 30 http://127.0.0.1:8080/ > benchmark/1-hello-world-roadrunner.txt

ab -c 1 -t 30 http://127.0.0.1:8081/
sleep 10
ab -c 1 -t 30 http://127.0.0.1:8081/ > benchmark/1-hello-world-apache-php.txt

ab -c 1 -t 30 http://127.0.0.1:8082/
sleep 10
ab -c 1 -t 30 http://127.0.0.1:8082/ > benchmark/1-hello-world-golang.txt

ab -c 1 -t 30 http://127.0.0.1:8085/de/blog/
sleep 10
ab -c 1 -t 30 http://127.0.0.1:8085/de/blog/ > benchmark/1-symfony-php-nginx.txt

ab -c 1 -t 30 http://127.0.0.1:8083/de/blog/
sleep 10
ab -c 1 -t 30 http://127.0.0.1:8083/de/blog/ > benchmark/1-symfony-roadrunner.txt


ab -c 4 -t 30 http://127.0.0.1:8080/
sleep 10
ab -c 4 -t 30 http://127.0.0.1:8080/ > benchmark/4-hello-world-roadrunner.txt

ab -c 4 -t 30 http://127.0.0.1:8081/
sleep 10
ab -c 4 -t 30 http://127.0.0.1:8081/ > benchmark/4-hello-world-apache-php.txt

ab -c 4 -t 30 http://127.0.0.1:8082/
sleep 10
ab -c 4 -t 30 http://127.0.0.1:8082/ > benchmark/4-hello-world-golang.txt

ab -c 4 -t 30 http://127.0.0.1:8085/de/blog/
sleep 10
ab -c 4 -t 30 http://127.0.0.1:8085/de/blog/ > benchmark/4-symfony-php-nginx.txt

ab -c 4 -t 30 http://127.0.0.1:8083/de/blog/
sleep 10
ab -c 4 -t 30 http://127.0.0.1:8083/de/blog/ > benchmark/4-symfony-roadrunner.txt


ab -c 32 -t 30 http://127.0.0.1:8080/
sleep 10
ab -c 32 -t 30 http://127.0.0.1:8080/ > benchmark/32-hello-world-roadrunner.txt

ab -c 32 -t 30 http://127.0.0.1:8081/
sleep 10
ab -c 32 -t 30 http://127.0.0.1:8081/ > benchmark/32-hello-world-apache-php.txt

ab -c 32 -t 30 http://127.0.0.1:8082/
sleep 10
ab -c 32 -t 30 http://127.0.0.1:8082/ > benchmark/32-hello-world-golang.txt

ab -c 32 -t 30 http://127.0.0.1:8085/de/blog/
sleep 10
ab -c 32 -t 30 http://127.0.0.1:8085/de/blog/ > benchmark/32-symfony-php-nginx.txt

ab -c 32 -t 30 http://127.0.0.1:8083/de/blog/
sleep 10
ab -c 32 -t 30 http://127.0.0.1:8083/de/blog/ > benchmark/32-symfony-roadrunner.txt


ab -c 1000 -t 30 http://127.0.0.1:8080/
sleep 10
ab -c 1000 -t 30 http://127.0.0.1:8080/ > benchmark/1k-hello-world-roadrunner.txt

ab -c 1000 -t 30 http://127.0.0.1:8081/
sleep 10
ab -c 1000 -t 30 http://127.0.0.1:8081/ > benchmark/1k-hello-world-apache-php.txt

ab -c 1000 -t 30 http://127.0.0.1:8082/
sleep 10
ab -c 1000 -t 30 http://127.0.0.1:8082/ > benchmark/1k-hello-world-golang.txt

ab -c 1000 -t 30 http://127.0.0.1:8085/de/blog/
sleep 10
ab -c 1000 -t 30 http://127.0.0.1:8085/de/blog/ > benchmark/1k-symfony-php-nginx.txt

ab -c 1000 -t 30 http://127.0.0.1:8083/de/blog/
sleep 10
ab -c 1000 -t 30 http://127.0.0.1:8083/de/blog/ > benchmark/1k-symfony-roadrunner.txt
```
