Este es el Cuarto Paso - Networking
------
## Task

## Introducción a Networking (Redes)

En docker, no nos preocupamos por gestionar los segmentos de red que podamos necesitar ya que es parte de su "trabajo"
por ende, solo nos preocupamos por indicarle que tipo de driver tendrá la red que creemos.

En docker existen 3 tipos de Redes:

- Host: el contenedor usará el mismo IP del servidor real que tengamos, para ello usa la interfaz del Host (ethX).
- Bridge: la red standard que usarán todos los contenedores, el segmento de red por defecto es 172.17.0.0/16.
- None: relacionado a la interfaz LO (localhost), se usa para pruebas puntuales como puede ser si el aplicativo inicia
de forma correcta.

Para listar las redes que tenemos ejecutamos: 

`docker network ls`{{execute}}

Con esta breve explicación vamos a crear una red de ejemplo: 

`docker network create red_tux`{{execute}}

Observación:

* Si queremos crear una red y asignarle un tipo especifico lo hacemos de la sgt manera:

`docker network create --driver bridge red_badtux`{{execute}}

* Si deseamos ver los segmentos de red que docker asigno a las redes: 

`docker network inspect red_tux`{{execute}}

* Si queremos eliminar una red:

`docker rm red_tux`{{execute}}

* Si queremos desconectar por temas de mantenimiento la red:

`docker network disconnect red_badtux`{{execute}}

y para volverlo a conectar usamos *connect*

* Si queremos crear una red y un segmento de red: 

docker network create --driver=bridge --subnet=192.168.2.0/24 --gateway=192.168.2.1 red_supertux
