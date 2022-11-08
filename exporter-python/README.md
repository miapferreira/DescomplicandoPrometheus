## Criando um exporter em Python

Detalhando um passo a passo para criar um exporter utilizando Python que vai nos trazer a métrica de quantas pessoas estão na espaço.
Para isso vamos bater na API [OpenNotify](http://api.open-notify.org/astros.json).


Crie o arquivo exporter.py conforme o exemplo deste [_arquivo_](https://github.com/miapferreira/prometheus/blob/master/exporter-python/exporter.py).

```bash
vim exporter.py
```

Ajuste a permissão de execução do arquivo e execute-o.

```bash
chmod +x exporter.py
python exporter.py
```

Abra um outro terminal execute o comando curl.

```bash
curl http://localhost:8899/metrics/
````

Verifique a saída que deverá ser algo como abaixo :D 

![log](prometheus/images/exporter_py_img.png)


## Adicionando nosso exporter em um container

Seguindo adiante vamos colocar nosso exporter para rodar num container Docker.

Caso ainda não tenha o docker instalado execute os passos abaixo para fazer a instalação.

```bash
curl -fsSL https://get.docker.com | bash
````

Crie o Dockerfile conforme o exemplo abaixo.

```bash
vim Dockerfile
````

```bash
FROM python:3.8-slim

LABEL maintainer Michel Aparecido Ferreira <seuemail@gmail.com.
LABEL description "Dockerfile para criar a imagem do nosso exporter"

WORKDIR /app
COPY . /app
RUN pip3 install -r requirements.txt

CMD python3 exporter.py
````

Crie o arquivo requirements.txt com as dependencias conforme abaixo

```bash
vim requirements.txt
````

```bash
requests
prometheus_client
```

Agora vamos realizar o build da nossa imagem.

```bash
docker build -t primeiro-exporter:0.1 .
````

Dando tudo certo com o build de nossa imagem podemos executar o container do nosso exporter em python.

```bash
docker run -p 8899:8899 --name primeiro-exporter -d primeiro-exporter:0.1
````

Cuncluído, podemos verificar se nosso exporter está rodando.

```bash
curl http://localhost:8899/metrics
```

