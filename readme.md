#COMANDOS PARA EXECUTAR O PROJETO

docker-compose up --build -d

cat loja_camisetas.sql | docker exec -i loja_camisetas_mysql /usr/bin/mysql -u root loja_camisetas
