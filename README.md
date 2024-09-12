# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados

## Instalación
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/WALC-1104/greetings
```

## Uso
package main

import (
	"fmt"
	"github.com/WALC-1104/greetings"
)

func main() {
	message, err := greetings.Hello("Test")

	if err != nil {
		fmt.println("Ocurrio un error:", err)
        return
	}

	fmt.Println(message)
}
```
Este ejemplo importa el paquete github.com/WALC-1104/greetings y llama la función Hello que retorna un saludo personalizado. si ocurre un error, se imprime un mensaje de error.