# goTuringMachine
Go implementation of a deterministic Turing machine, given its rules as a JSON file.

### Compilación:
Para compilar, primero ubíquese en el directorio del
programa:
```
    cd /(the path to the project directory)/goTuringMachine
```
Para compilar, ejecute los siguientes comandos:
```
    export GOPATH=$(pwd)
    export GOBIN=$(pwd)/bin
    go install ./...
```
lo cual, genera el binario "main", el cual contiene el
programa compilado, éste se ejecuta mediante:
```
    ./bin/main
```

### Ejecución:
El programa preguntará primeramente por un nombre de
archivo, en el directorio donde se ejecuta:
```
    Archivo con la descripcion de la MT : [su archivo JSON de la descripcion va aquí]
```
seguido de la cadena que se desea dar a la Máquina de Turing
como entrada:
```
    Inserte la cadena de entrada : [su cadena va aquí]
```

### Ejecución ejercicio 2:
Se adjunta la especificación de la máquina de Turing para el ejercicio 2, con un archivo llamado 'desc2.json', el cual, una vez compilado el programa, puede ser ejecutado de la siguiente manera :
```
    Archivo con la descripcion de la MT : desc2.json
```
seguido de la cadena que se desea dar a la Máquina de Turing
como entrada (un ejemplo es el siguiente):
```
    Inserte la cadena de entrada : 11x11=1111
```

### Autores:
Arteaga Vázquez Alan Ernesto

Figueroa Sandoval Gerardo Emiliano
