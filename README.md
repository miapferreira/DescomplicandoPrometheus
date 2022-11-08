# Prometheus

Este projeto tem como objetivo documentar e compartilhar o conhecimento adquiridido durante os meus estudos sobre o prometheus.


## Referências

 - [Treinamento LinuxTips Descomplicando o Promethueus](https://github.com/badtuxx/DescomplicandoPrometheus)
 - [Documentação oficial Prometheus](https://prometheus.io/)


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

Verifiquei se o binário está funcionando.

```bash
prometheus --version

prometheus, version 2.38.0 (branch: HEAD, revision: 818d6e60888b2a3ea363aee8a9828c7bafd73699)
  build user:       root@e6b781f65453
  build date:       20220816-13:23:14
  go version:       go1.18.5
  platform:         linux/amd64
```

Crie os seguintes diretórios de necessários para configuração do prometheus

```bash
sudo mkdir /etc/prometheus
sudo mkdir /var/lib/prometheus
sudo mkdir /var/log/prometheus
````

Crie um grupo e um usuário para o Prometheus.

```bash
sudo addgroup --system prometheus
sudo adduser --shell /sbin/nologin --system --group prometheus
```

Mude a permissão para que o usário prometheus criado no passo anterior seja o dono dos diretórios.

```bash
sudo chown -R prometheus:prometheus /var/log/prometheus
sudo chown -R prometheus:prometheus /etc/prometheus
sudo chown -R prometheus:prometheus /var/lib/prometheus
sudo chown -R prometheus:prometheus /usr/local/bin/prometheus
sudo chown -R prometheus:prometheus /usr/local/bin/promtool
````


Seguindo, vamos precisar mover os consoles, console_libraries e o arquivo prometheus.yml para o diretório de configuração do Prometheus que acabamos de criar no passo acima

```bash
sudo mv prometheus-2.38.0.linux-amd64/prometheus.yml /etc/prometheus/prometheus.yml
sudo mv prometheus-2.38.0.linux-amd64/consoles /etc/prometheus
sudo mv prometheus-2.38.0.linux-amd64/console_libraries /etc/prometheus
```

Edite o arquivo de configuração prometheus.yml conforme este [_repositório_](https://github.com/miapferreira/prometheus/blob/master/conf/prometheus.yml)

```bash
sudo vim /etc/prometheus/prometheus.yml
```

Será necessário fazer com que o Prometheus seja um serviço em nossa máquina, para isso precisamos criar o arquivo chamado prometheus.service o exemplo deste [_repositório_](https://github.com/miapferreira/prometheus/blob/master/conf/prometheus.service)

```bash
sudo vim /etc/systemd/system/prometheus.service
```

Recarregue o systemd e inicie o serviço do prometheus.

```bash
sudo systemctl daemon-reload
sudo systemctl start prometheus
````

Habilite o serviço prometheus para que seja iniciado automaticamente ao iniciar o sistema.

```bash
sudo systemctl enable prometheus
````

Para garantir, verifique o status do serviço

```bash
sudo systemctl status prometheus
````

Finalizando todo processo de instalação, acesse a interface web do Prometheus em seu navegador.

```bash
http://localhost:9090
````
Por fim, teremos completado nossa instalação se seu navegador mostrar a imagem abaixo :D 

![Prometheus instalado no Linux](images/prometheus_interface_web.png)


# Tipos de dados utilizados pelo Prometheus
## Gauge

O Gauge trabalha com variacao de dados, como por exemplo temperatura, consumo de cpu, memória, etc..

## Counter 

Sao dados que vao ser incrementados ao longo do tempo, como por exemplo a quantidade de requisicoes com sucesso 
uma aplicacao recebeu (HTTP status code 200, 4xx, 5xx)

## Histogram

Utiliza buckets que já sao pré-definidos com seu valor, por exemplo,
quero consultar o tempo de que cada requisicao durou
- 0,5 segundos -> 100 requests 
- 1 segundos -> 150 requisicao
- 2 segundos -> 37 requests

## Summary

O sumary funciona como histogram, a diferenca basicamente é que os buckets sao chamados 
de quantiles eu nao vou ter a mesma flexibilidade de consulta no 
momento da query, ele já vai me retornar os valores exatos, ou seja, o seu resumo.
