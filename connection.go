package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var data = map[string]string{} // cria uma 'map' vazio
func browserHeader(conn net.Conn, str string) {
	//cabeçalho necessário para enviar uma mensagem ao Browser.
	length := fmt.Sprint("Content-Length: ", len(str), "\r\n")
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte(length))
	conn.Write([]byte("Content-Type: text/html\r\n\r\n"))
	conn.Write([]byte(str))
}

func response(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	rote := strings.Split(scanner.Text(), " ")[1]
	path := strings.Split(rote, "/") //separa as rotas '/'
	fmt.Println(path)
	switch path[1] {
	case "le":
		fmt.Println("Requisição para leitura.")
		resp := fmt.Sprintln(path[2], "tem o valor:", data[path[2]])
		if len(path[2]) <= 0 {
			browserHeader(conn, "Campo vazio!")
		}
		browserHeader(conn, resp)
	case "escreve":
		fmt.Println("Requisição para escrita:")
		data[path[2]] = path[3]
		browserHeader(conn, "O valor foi escrito!")
	case "delete":
		fmt.Println("Requisição para deletar:")
		delete(data, path[2])
		browserHeader(conn, "O valor foi deletado com sucesso!")
	}
	conn.Close()
}

/*
http://localhost:3000/escreve/campo4/golang-web (insere a frase 'golang-web' na variavel 'campo4')
http://localhost:3000/le/campo4/		        (exibe o conteúdo da variavel 'campo4')
http://localhost:3000/delete/campo4/			(deleta o conteúdo do 'campo4')
*/
