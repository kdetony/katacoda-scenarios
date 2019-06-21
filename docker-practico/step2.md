Este es el Segundo Paso.

## Task

Vamos a buscar imagenes, para ellos usamos el siguiente comando  **docker search IMAGE** por ejemplo:

`docker search nginx`{{execute}}

Por defecto Docker Engine buscar desde un registry Public (dockerhub) las imagenes que necesita.
A su vez podemos crearnos una cuenta en Docker Hub, crear nuestro repository Publico/Privado y 
subir/descargar nuestras propias imagenes.

Ahora ejecutemos el comando **docker pull IMAGE** para descargar una imagen 

`docker pull nginx`{{execute}}

De esta forma descargamos la imagen de Nginx en nuestro Host, pero es bueno comentar lo siguiente, 
la imagenes deben identificarse previamente, es decir, yo puedo descargar versiones especificas, dependiendo
de la necesidad, a su vez si no se especifica, se sobrentiende que es la version *latest* 
