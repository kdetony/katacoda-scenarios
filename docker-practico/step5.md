Este es el Quinto Paso(Imagenes)
----------
## Task

Debemos recordar que los datos para los contenedores son efimeros, es decir, que si los borramos toda la información
que podamos haber ingresado al contendor se perderá. Para evitar estos problemas recurrimos a los volumenes, los cuales
pueden ser: 

- Volume: es la manera sencilla, simple y predefinida para almacenar los datos para un contenedor, usará el espacio del Host y en */var/lib/docker/volumes* creara una carpeta para cada volumen creado. Esto es similar a tener un datastore.

``docker volume create VG1``{{execute}}

Para asociarlo a un contenedor: 

``docker run -dit -v VG1:/usr/local/apache2/htdocs/ httpd``{{execute}}

- Volume Bind: es una manera de "asociar" una carpeta del Host y mapearla como una carpeta dentro de un contenedor.

Para crear un volume de este tipo ejecutaremos: 

``docker run -rm -dit -v /tmp/webserver:/usr/local/apache2/htdocs/ httpd``{{execute}}


- TMPFS: (temporal file system) es una manera de montar carpetas temporales en un contenedor. Usan la RAM del Host y su contenido desaparecerá al parar el contenedor.

OBS.
En caso de tener poca RAM, los ficheros se parasán al SWAP del equipo real.

Para crear un volume de este tipo ejecutaremos:

`docker run -dit --tmpfs /usr/local/apache2/htdocs/ httpd `{{execute}}
