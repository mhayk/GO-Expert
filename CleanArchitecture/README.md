# Clean Architecture

* Termo criado por Robert C. Martin (Uncle Bob) em 2012.
* É uma forma de organizar o código de forma que ele seja independente de frameworks, UI, banco de dados, etc.
* O objetivo é criar um código que seja fácil de manter, testar e evoluir.
* O código é dividido em camadas, onde cada camada tem uma responsabilidade.
* As camadas são independentes e se comunicam através de interfaces.
* A camada mais interna é a mais importante, pois é a que contém as regras de negócio.
* A camada mais externa é a que se comunica com o mundo externo (UI, banco de dados, etc).
* A camada mais externa depende das camadas mais internas, mas não o contrário.
* A camada mais interna não sabe da existência das camadas mais externas.
* A camada mais interna é a mais fácil de testar, pois não depende de nada.
* A camada mais externa é a mais difícil de testar, pois depende de muitas coisas.
* A camada mais externa é a que mais muda, pois depende de frameworks, UI, banco de dados, etc.
* É orientada a casos de uso, onde cada caso de uso é uma funcionalidade do sistema.

## Use Cases
* Casos de uso são as funcionalidades do sistema.
* Cada caso de uso é uma classe que implementa uma interface.
* Itenção
* Clareza de cada comportamento do software
* Detalhes não devem impactar nas regras do negócio
* Frameworks, banco de dados, apis, não devem impactar nas regras de negócio

### Use Cases vs SRP
* Temos a tendência de "reaproveitar" use cases por serem muito parecidos.

## DTO (Data Transfer Object)
* São objetos que carregam dados entre as camadas.
* Trafegar dados entre os limites arquiteturais
* Objeto anêmico, sem comportamento
* Contém dados (Input ou Output)
* Não possui regras de negócio
* Não possui comportamento
* NÃO faz nada!
* API -> CONTROLLER -> USE CASE -> ENTITY
* Controller cria um DTO com os dados recebidos e envia para o Use Case
* Use Case executa seu fluxo, pega o resultado, cria um DTO para output e retetorna para o Controller