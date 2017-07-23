package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func tratar(s string, e error) {
	if e != nil {
		log.Fatalln(s + e.Error())
	}
}

func main() {
	var porta string
	fmt.Print("Digite o numero da porta a ser aberta: ")
	fmt.Scan(&porta)
	li, err := net.Listen("tcp", ":"+porta)
	tratar("Erro ao abrir servidor: ", err)
	defer li.Close()

	fmt.Printf("%v\n", `Pressione as teclas "ctrl + c" para fechar o servidor.`)
	fmt.Println("Servidor aguardando conexão...")

	for {
		conn, err := li.Accept()
		tratar("Erro ao receber conexao: ", err)
		c := make(chan bool)

		fmt.Printf("%v\n", "Cliente conectou-se.")
		fmt.Fprintf(conn, "%v\n", `Digite "sair" para fechar a conexao.`)
		go ler(&conn, c)
		go escrever(&conn, c)
	}
}

func ler(conn *net.Conn, c chan bool) {
	scanner := bufio.NewScanner(*conn)
	dc := func() {
		fmt.Println("Cliente desconectou-se.")
		c <- true
		close(c)
	}

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "sair" {
			break
		} else {
			fmt.Printf("%v\n", "Cliente: "+ln)
		}
	}
	dc()
	defer (*conn).Close()
}

func escrever(conn *net.Conn, c chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	saiu := false

	go func() {
		<-c
		saiu = true
	}()
	for scanner.Scan() && !saiu {
		ln := scanner.Text()
		fmt.Fprintf(*conn, "%v\n", "Servidor: "+ln)
	}
}
