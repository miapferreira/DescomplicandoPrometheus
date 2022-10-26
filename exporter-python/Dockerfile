FROM python:3.8-slim

LABEL maintainer Michel Aparecido Ferreira <mi.apferreira@gmail.com.
LABEL description "Dockerfile para criar a imagem do nosso exporter"

WORKDIR /app
COPY . /app
RUN pip3 install -r requirements.txt

CMD python3 exporter.py
