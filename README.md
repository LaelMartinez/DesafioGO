Nome do Projeto
Descrição
Este é um aplicativo Go simples que lê um arquivo CSV contendo informações sobre usuários (Nome, Idade, Pontuação) e gera dois arquivos de saída ordenados, um por nome e outro por idade.

Funcionalidades
Leitura de CSV
Ordenação por Nome
Ordenação por Idade
Geração de dois arquivos de saída ordenados
Pré-requisitos
Go 

Instalação
Clone o repositório:

bash
git clone [https://github.com/seu-usuario/seu-projeto.git](https://github.com/LaelMartinez/DesafioGO)
cd seu-projeto

Instale as dependências:
bash
go mod tidy

Como Usar
Execute o aplicativo com o comando:
bash

go run main.go <arquivo-origem.csv> <arquivo-destino.csv>
Substitua arquivo-origem.csv pelo caminho do seu arquivo CSV.

Estrutura do CSV
O CSV deve ter as seguintes colunas:

Nome
Idade
Pontuação
Exemplo:

csv
Copy code
Nome,Idade,Pontuação
Alice,30,90
Bob,25,85
Charlie,35,92

Contribuição
Fiquem a vontade para contribuir com este projeto.
