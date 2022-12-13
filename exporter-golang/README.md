
# Criando um exporter em Go

Detalhando um passo a passo para criar um exporter utilizando Golang para nos trazer métricas referente ao consumo de  memória.

Crie o arquivo exporter.go conforme o exemplo deste [_arquivo_](https://github.com/miapferreira/prometheus/blob/master/exporter-golang/exporter.go).

```bash
vim exporter.go
```

Instale as seguintes bibliotecas necessárias para nosso código

```bash
go mod init exporter.go
go mod tidy
````

Em seguida já podemos compilar nosso código

```bash
go build exporter.go
````
Perceba que será gerado um binário chamado exporter. Para testa-lo basta executa-lo e em seguida verificar as métricas conforme a porta que
configuramos em nosso exporter.go, ou seja, na porta 7788.

```bash
./exporter
curl http://localhost:7788/metrics
````
A saída deve ser algo semelhante a imagem abaixo contendo as métricas coletadas

![log](https://github.com/miapferreira/prometheus/blob/master/images/go.png)

## Adicionando o exporter em um container

Crie o arquivo Dockerfile conforme abaixo.

```bash
vim Dockerfile
````

```bash
FROM golang:1.19.0-alpine3.16 AS buildando

WORKDIR /app
COPY . /app

RUN go build exporter.go



FROM alpine:3.16

COPY --from=buildando /app/exporter /app/exporter
EXPOSE 7788
WORKDIR /app
CMD ["./exporter"]
````

Realize o buil da imagem

```bash
docker build -t segundo-exporter:1.0 .
````

Dando tudo certo com o build de nossa imagem podemos executar o container do nosso exporter em Go.

```bash
docker run -d --name segundo-exporter -p 7788:7788 segundo-exporter:1.0
````
verifique se o container está sendo executado.

```bash
docker ps 
````
```bash
CONTAINER ID   IMAGE                  COMMAND        CREATED         STATUS         PORTS                                       NAMES
163d10dfd379   segundo-exporter:1.0   "./exporter"   6 seconds ago   Up 4 seconds   0.0.0.0:7788->7788/tcp, :::7788->7788/tcp   second-exporter
````

Acesse as métricas do nosso exporter, dando tudo certo o resultado será conforme a ![imagem](https://github.com/miapferreira/prometheus/blob/master/images/go.png)

Também é necessário adicionar o novo Target para que o prometheus possa coletar as métricas, para isso adicione as informacoes abaixo no arquivo 
[_prometheus.yml_](https://github.com/miapferreira/prometheus/blob/master/conf/prometheus.yml)

```bash
- job_name: "Segundo Exporter Golang"
    static_configs:
      - targets: ["localhost:7788"]
```
Reinicie o servico do prometheus e acompanhe pelo navegador o novo targer registrado

```bash
systemctl restart prometheus
http://localhost:9090
```