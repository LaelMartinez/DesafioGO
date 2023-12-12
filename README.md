<h1>Leitura de geração de arquivos .csv usando GoLang</h1>
  
<h2>Descrição</h2>
Este é um aplicativo Go simples que lê um arquivo CSV contendo informações sobre usuários (Nome, Idade, Pontuação) e gera dois arquivos de saída ordenados, um por nome e outro por idade.

<h2>Funcionalidades</h2>
<h2>Leitura de CSV</h2>
  
Ordenação por Nome
Ordenação por Idade
Geração de dois arquivos de saída ordenados

<h2>Pré-requisitos</h2>
Go 

<h2>Instalação</h2>
Clone o repositório:

bash
git clone [https://github.com/seu-usuario/seu-projeto.git](https://github.com/LaelMartinez/DesafioGO)
cd seu-projeto

<h2>Instale as dependências:</h2>
bash
go mod tidy

<h2>Como Usar</h2>
Execute o aplicativo com o comando:
bash

go run main.go <arquivo-origem.csv> <arquivo-destino.csv>
Substitua arquivo-origem.csv pelo caminho do seu arquivo CSV.

<h2>Estrutura do CSV</h2>
O CSV deve ter as seguintes colunas:

Nome
Idade
Pontuação
Exemplo:

csv
Nome,Idade,Pontuação
Alice,30,90
Bob,25,85
Charlie,35,92

<h2>Contribuição</h2>
Fiquem a vontade para contribuir com este projeto.
