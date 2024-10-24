package proxy

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProxyRequest redirige la solicitud al microservicio correspondiente
func ProxyRequest(c *gin.Context, targetURL string) {
	// Crear un cliente HTTP
	client := &http.Client{}
	reqBody, _ := ioutil.ReadAll(c.Request.Body)

	// Crear una nueva solicitud HTTP
	req, err := http.NewRequest(c.Request.Method, targetURL, bytes.NewBuffer(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud"})
		return
	}

	// Copiar los encabezados originales de la solicitud
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Comprobar si el username existe en el contexto y agregarlo al encabezado si existe
	if username, exists := c.Get("username"); exists {
		req.Header.Set("X-Auth-Username", username.(string)) // Añadir el username al encabezado
	}

	// Realizar la solicitud al microservicio
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con el microservicio"})
		return
	}

	// Leer la respuesta del microservicio y devolverla al cliente
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
