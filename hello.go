package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

//import "reflect"

func main() {
	//É possível declarar as variáveis com var e tipar as mesmas
	/*var nome string = "Daniel"
	var idade int
	var versao float32 = 1.1
	//fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))*/

	/*if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não conheço esse comando")
	}*/

	//fmt.Scanf("%d", &comando)
	//fmt.Println("O endereço da minha variável comando é", &comando)

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	//nome := "Daniel"
	nome, _ := devolveNomeEIdade()
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	pulaLinha()
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func devolveNomeEIdade() (string, int) {
	nome := "Daniel"
	idade := 31

	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	/*var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"*/

	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	for j := 0; j < monitoramentos; j++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		pulaLinha()
	}

	pulaLinha()
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") //bytes
	//arquivo, err := ioutil.ReadFile("sites.txt") //buffer de bytes, string(arquivo) //Pacote io.ioutil

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		//fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		registraLog(site, true)
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		registraLog(site, false)
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func registraLog(site string, status bool) {
	arquivo, err := os.Open("log.txt")
	os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	fmt.Println(arquivo)
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

func pulaLinha() {
	fmt.Println("")
}

//Funções de teste de Slice
func exibeNomes() {
	nomes := []string{"Daniel", "Mendes", "Melo"}
	printSliceInfo(nomes)
	nomes = append(nomes, "Sousa")
	printSliceInfo(nomes)
}

func printSliceInfo(nomes []string) {
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
