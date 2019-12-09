# SwaggerUI for Openfaas

Simple utility to server swagger static files and reverse proxy requests openfaas platform.

Provide the environment variable "openfaas_gateway"

This utility requires the openfaas function mentioned [here](https://github.com/Optum/faas-swagger/tree/master/swagger-as-function) to work.

```
docker run -it -e openfaas_gateway=http://gateway:8080 -p 8080:8080 murugappans/swaggerui-openfaas
```