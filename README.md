# labs-one
Projeto desenvolvido para o primeiro laboratório da pós go

- criar a imagem
docker build -t cloudrun .

- rodar imagem
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e URL_CEP="https://viacep.com.br/ws" \
  -e URL_TEMPO="https://api.weatherapi.com/v1" \
  -e API_KEY_TEMPO="3baa5b20172b4baf91c185158251003" \
  cloudrun


https://labs-one-1008210695250.us-central1.run.app