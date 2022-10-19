# prometheus

## Instalando e configurando o Prometheus no Linux

1. Realize o download da versão mais recente do Prometheus utilizando o comando abaixo, aqui vamos utilizar a versão [2.38.0]
curl -LO https://github.com/prometheus/prometheus/releases/download/v2.38.0/prometheus-2.38.0.linux-amd64.tar.gz

2. Após o download, descompactue o arquivo
tar -xvf prometheus-2.38.0.linux-amd64.tar.gz

3. Como próximo passo, vamos remover os binários para o diretório /usr/local/bin 
sudo mv prometheus-2.38.0.linux-amd64/prometheus /usr/local/bin/prometheus
sudo mv prometheus-2.38.0.linux-amd64/promtool /usr/local/bin/promtool
