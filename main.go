package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Pessoa struct {
	Nome      string
	Idade     int
	Pontuacao int
}

type Pessoas []Pessoa

func (p Pessoas) Len() int { return len(p) }

func (p Pessoas) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pessoas) Less(i, j int) bool { return p[i].Nome < p[j].Nome }

type PorIdade struct{ Pessoas }

func (p PorIdade) Less(i, j int) bool { return p.Pessoas[i].Idade < p.Pessoas[j].Idade }

func LerArquivo(nomeArquivo string) (Pessoas, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	leitorCSV := csv.NewReader(arquivo)
	linhas, err := leitorCSV.ReadAll()
	if err != nil {
		return nil, err
	}

	var pessoas Pessoas
	for _, linha := range linhas {
		idade, _ := strconv.Atoi(linha[1])
		pontuacao, _ := strconv.Atoi(linha[2])

		p := Pessoa{
			Nome:      linha[0],
			Idade:     idade,
			Pontuacao: pontuacao,
		}
		pessoas = append(pessoas, p)
	}

	return pessoas, nil
}

func EscreverArquivo(nomeArquivo string, pessoas Pessoas) error {
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	escritorCSV := csv.NewWriter(arquivo)
	defer escritorCSV.Flush()

	for _, pessoa := range pessoas {
		linha := []string{pessoa.Nome, strconv.Itoa(pessoa.Idade), strconv.Itoa(pessoa.Pontuacao)}
		err := escritorCSV.Write(linha)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <arquivo-origem.csv> <arquivo-destino.csv>")
		os.Exit(1)
	}

	arquivoOrigem := os.Args[1]
	arquivoDestino := os.Args[2]

	pessoas, err := LerArquivo(arquivoOrigem)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	sort.Sort(pessoas)
	err = EscreverArquivo(arquivoDestino+"_ordenado_por_nome.csv", pessoas)
	if err != nil {
		fmt.Printf("Erro ao escrever o arquivo ordenado por nome: %v\n", err)
		os.Exit(1)
	}

	sort.Sort(PorIdade{pessoas})
	err = EscreverArquivo(arquivoDestino+"_ordenado_por_idade.csv", pessoas)
	if err != nil {
		fmt.Printf("Erro ao escrever o arquivo ordenado por idade: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Processamento concluído com sucesso.")
}
