docker run --rm --name mysqld\
  -p 3306:3306\
  -e MYSQL_ROOT_HOST=%\
  -e MYSQL_ALLOW_EMPTY_PASSWORD=yes\
  mysql/mysql-server:latest

mysqldump -u root -h 127.0.0.1 jide > jide.sql

---

testing notes
$ go test ./... -coverprofile=cover.out -coverpkg=./...
$ go tool cover -html=cover.out
