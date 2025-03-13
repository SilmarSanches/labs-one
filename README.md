# Labs-One

Projeto de conclusão de pós-graduação

## Indice
1. [Build](#build)
2. [Dockerfile](#dockerfile)
3. [Testes](#testes)
4. [Swagger](#swagger)


## Build
- criar a imagem:
```bash
docker build -t cloudrun .
```

## Dockerfile
- rodar imagem:
```bash
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e URL_CEP="https://viacep.com.br/ws" \
  -e URL_TEMPO="https://api.weatherapi.com/v1" \
  -e API_KEY_TEMPO="3baa5b20172b4baf91c185158251003" \
  cloudrun
```

## Testes
- testes unitários podem ser executados com o comando:
```bash
go test ./...
```
- testes de integração podem ser executados com o arquivo testes.http que está na pasta /app. Neste arquivos temos testes que podem ser executados locais na porta 8080 e testes que podem ser executados na google cloud run através do endereço: 

```bash
https://labs-one-l5e2u7s2ca-uc.a.run.app
```


## Swagger
- o swagger pode ser acessado localment na porta 8080 ou na google cloud run atraves dos endereços abaixo:

```bash
http://localhost:8080/swagger
```

```bash
https://labs-one-l5e2u7s2ca-uc.a.run.app/swagger/index.html#/
```
