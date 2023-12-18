# Prueba técnica Golang Amaris

## Backend

### Arquitectura hexagonal

![image](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/72397b67-fa49-48ac-9a48-df28671bb1a0)

Se implementa la arquitectura hexagonal, la organización de las capas del proyecto queda de la siguiente manera:

![Captura de pantalla 2023-12-17 230303](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/c0096e1f-d011-4ab5-a870-8154709a7652)

### Cobertura de pruebas unitarias mayor al 80%

![Captura de pantalla 2023-12-17 234922](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/de5747d8-6bf0-43e0-9b9b-a44f2ff88296)

### Patrón CQRS

Se implementa el patrón CQRS en donde están separadas la lógica de lectura y escritura de datos en dos DB DynamoDB, para mayor escalabilidad cada una tiene su propia DB, una en local para las escrituras y otra en cloud para las lecturas.

![Captura de pantalla 2023-12-17 235038](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/fe50ef79-d861-4ebb-ad7a-e605d378057a)

![Captura de pantalla 2023-12-17 231646](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/0b65a699-66e9-4176-8efa-b652697bb96b)

![Captura de pantalla 2023-12-17 231921](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/458e8ebc-0a2f-49df-9c0e-6423bd03ca42)


### Broker de comunicación

Para mantener la información actualizada y consistente en ambas DB se hizo una implementación en el código usando go rutinas de manera asyncrona, para que cada vez que se actualice la DB de escritura se envíe asíncronamente una actualización a la DB de lectura.

![Captura de pantalla 2023-12-17 235941](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/ec6434aa-27ad-4947-a29d-62e86ec7c0c4)


### Curls de endpoints

* Para la consulta de puntos por id:

CURL: ``` curl --location 'http://localhost:8080/id/1' ```

![Captura de pantalla 2023-12-17 231353](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/e93d42cb-7e5b-4a5f-b10f-b1cb9c727b16)


* Para la actualización de puntos:

CURL: ``` curl --location --request POST 'http://localhost:8080/points/200/id/1' ```  

![Captura de pantalla 2023-12-17 231442](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/5d4c2fde-44d3-4fdc-a5ea-49453354f21f)

### Desplegar el proyecto

1. Se debe tener instalado previamente Docker y node.js en el computador.
2. Desde la consola de admin en windows o con un usuario sudo en linux/mac ejectutar el siguiente comando para desplegar el BynamoDB en local: ``` docker run -p 8000:8000 amazon/dynamodb-local  ```

Deberá aparecer una imagen en Docker como la siguiente:

![Captura de pantalla 2023-12-17 233615](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/c0fe913b-660d-46a5-a624-e04a9ad37914)

4. Si desea ver el admin de DynamoDB en local deberá ejecutar los siguientes comandos en una consola de admin: ``` npm install -g dynamodb-admin ```, ``` set DYNAMO_ENDPOINT=http://localhost:8000 ``` y  ``` dynamodb-admin  ```
5. Luego al entrar al link ``` http://localhost:8001/ ``` podrá ver el admin de Dynamo

![Captura de pantalla 2023-12-17 233725](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/0e645d5d-f46a-4bce-9929-b377814083d4)

6. Luego en el proyecto hay que ir a la carpeta y el archivo backend/config/config.yaml y cambiar las variables secret y key de AWS por las que se encuentran en el email de respuesta de esta prueba.

![Captura de pantalla 2023-12-17 234013](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/8bba36ff-9314-4cb3-a9bd-8feba0930ef4)

8. También en todos los archivos de test hay que modificar la ubicación del archivo yaml por la ubicación física del archivo en el computador:

![Captura de pantalla 2023-12-17 235424](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/1ed85bd8-c39b-43b4-9030-75b8c0de423d)

9. Una vez hecho todas las configuraciones puede desplegar el proyecto ubicandose en la carpeta raíz "backend" y desde la consola ejecutando los siguientes comandos: ``` go mod tidy ```, ``` go run main.go ```
10. Ya podrá consumir los endpoints, están descritos en la sección anterior "Curls de endpoints"

# Frontend

El frontend fue desarrollado en React, a continuación detallo su funcionamiento:

1. Página principal y consulta de puntos por id:

![Captura de pantalla 2023-12-17 232513](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/eed43ff5-f777-40fd-bfe9-6083a4659b50)

![Captura de pantalla 2023-12-17 232431](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/2b2037dc-db5e-4a82-92c3-057e5d5daccd)

2. Summar puntos:

![Captura de pantalla 2023-12-17 232712](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/6bc75688-0c01-4404-a5f0-dffe8f6cb77c)

![Captura de pantalla 2023-12-17 232743](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/7647f6a9-f2e7-4df3-abfd-8be2b5e2b66f)


3. Restar puntos:

![Captura de pantalla 2023-12-17 232534](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/52e42c73-1b0a-44e2-9a97-0bdf5a495185)

![Captura de pantalla 2023-12-17 232628](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/5df1b94d-841a-4e34-9e92-222c8546e89c)

4. Cuando los puntos dan un total de 500 se gana un premio:

![Captura de pantalla 2023-12-17 232832](https://github.com/diazclaudia/Amaris_Golang_test/assets/16843197/0536c6be-9b50-47fa-9e9e-d98af92721ee)


## Despliegue

1. Debe tener instalado node.js
2. Debe ubicarse en la raíz del proyecto frontend/my-app y ejecutar el comando: ``` npm install ```
3. Luego ejecutar el comando para desplegar: ``` npm start ```
4. Abrir el link: ``` http://localhost:3000/ ```









