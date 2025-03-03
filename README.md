docker pull mysql:latest
docker run --name bdorm-container -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=mydatabase -p 3306:3306 -d mysql:latest
