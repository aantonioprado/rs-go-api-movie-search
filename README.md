# API Movie Search

API desenvolvida durante o curso de **Go da Rocketseat**, com o objetivo de praticar conceitos de rotas, requisições HTTP e consumo de APIs externas.

O projeto faz integração com a **OMDb API**, permitindo buscar filmes por título e obter detalhes de cada um deles.

## 🚀 Tecnologias

- **Go** (Golang)  
- **OMDb API** (The Open Movie Database)

## ⚙️ Funcionalidades

- Buscar filmes por título  
- Obter detalhes de um filme específico

## 💡 Sobre o projeto

Esse projeto foi desenvolvido como parte do **curso de Go da Rocketseat**, com foco em entender:
- Estrutura básica de uma aplicação Go  
- Criação de rotas HTTP  
- Consumo de APIs externas  
- Manipulação de respostas JSON  

## ▶️ Como executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/aantonioprado/rs-go-api-movie-search
   cd rs-go-api-movie-search
   ```

2. Configure sua chave da OMDb:
   ```bash
   export APIKEY_OMDB=<key>
   ```

3. Execute o projeto:
   ```bash
   go run .
   ```

4. Acesse no navegador:
   ```
   http://localhost:8080/?s=blade
   ```
