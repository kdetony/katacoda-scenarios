Este es el Tercer Paso - Contenedores.
-------
## Task

En el anterior ejemplo vimos 2 comandos importantes: *search* y *pull* (buscar/descargar imagenes respectivamente)
Ahora vamos a iniciar un contenedor, debemos recordar que, para que el contenedor pueda inicializar, este debe
siempre contar con una imagen asociada, dicho en otras palabras **la imagen es la punta de lanza para el contenedor**.

Vamos ahora usar el comando **docker run** el cual nos permite iniciar contendores:

`docker run nginx`{{execute}}

Podemos visualizar que el contenedor como tal inicia y por ende tenemos la salida del mismo en pantalla. Algo molesto cierto?
debido a que no podemos interactuar con el bash para ejecutar otros comandos, vamos a cancelar el contenedor usando:
```markdown
CTRL+C
```
Vamos a iniciar de una forma distinta el contendor, para ello ejecutaremos:

`docker run -dit nginx`{{execute}}

Veremos que al final de la ejecución del comando se obtiene un codigo similar a este: 
```markdown
09asn8jnn0823inoinfoi
```

Ahora ejecutemos este comando: 

`docker ps -a`{{execute}}

Este comando es muy importante porque nos muestra la sgt información:

* ID_Contenedor
* Hora de creación
* Nombre 
* Puertos 

OBS.
* Si queremos detener un contendor: 
```markdown
docker stop ID_CONTAINER
```

* Si queremos borrar un contenedor:
```markdowon
docker rm ID_CONTAINER
```
