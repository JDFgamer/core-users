# Usa una imagen base de Go
FROM golang:1.23-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de tu aplicación al contenedor
COPY . .

# Copia los archivos de configuración YAML al contenedor
COPY ./config /app/cmd/config

# Establece la variable de entorno para el entorno local
ENV CONFIG_ENV=local

# Establece el directorio de trabajo a la carpeta que contiene el archivo main.go
WORKDIR /app/cmd

# Descarga las dependencias necesarias para tu aplicación
RUN go mod tidy

# Construye la aplicación Go
RUN go build -o /app/users .

# Expone el puerto que tu aplicación va a usar
EXPOSE 8080

# Comando para ejecutar la aplicación Go cuando el contenedor se inicie
CMD ["/app/users"]
