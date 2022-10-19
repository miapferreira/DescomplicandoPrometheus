## Instalação do prometheus no linux

Realize o download da versão mais recente do Prometheus utilizando o comando abaixo, aqui vamos utilizar a versão [2.38.0].

```bash
curl -LO https://github.com/prometheus/prometheus/releases/download/v2.38.0/prometheus-2.38.0.linux-amd64.tar.gz
```

Após o download, descompactue o arquivo:

```bash
tar -xvf prometheus-2.38.0.linux-amd64.tar.gz
```
    
Como próximo passo, vamos remover os binários para o diretório /usr/local/bin

```bash
sudo mv prometheus-2.38.0.linux-amd64/prometheus /usr/local/bin/prometheus
sudo mv prometheus-2.38.0.linux-amd64/promtool /usr/local/bin/promtool
```