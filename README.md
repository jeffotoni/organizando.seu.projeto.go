# Organizando seu projeto Go.


Vamos criar pacotes Go, vamos tentar entender um pouco quando estamos criando projetos em Go.


Tudo inicia em Go mod, o que temos que saber e como usa-lo.

 - go mod init 
 	Cria um novo módulo, inicializando o go.mod

 - go build
 	Testa outros comandos de construção de pacote adicionam novas dependências go.mod conforme necessário.

 - go list -m all 
 	Imprime as dependências do módulo atual.

 - go get 
 	Altera a versão necessária de uma dependência (ou adiciona uma nova dependência).

 - go mod tidy 
 	Remove dependências não utilizadas.


Tudo começa criando pacotes, entendendo como atualiza-los e como organinza-los.

Antigamente não tinhamos um gerenciador de dependências então era usado outras ferramentas pela comunidade.

Em outras langs temos:

 - Composer do PHP
 - yarn/npm para JavaScript
 - pip no Python
 - Maven/Gradle no Java

 Em Go quando não tinhamos o Go mod usavamos:

  - Dep
  - Glide

A partir da versão go1.11 e go1.12 temos o go mod de forma preliminar e na versão go1.13 temos o go mod como padrão para todo desenvolvimento

Sabemos que não tem nada tão simples assim, ao escrever um aplicativo por exemplo por mais simples que seja dificilmente não iremos utilizar uma lib de terceiro.
Em um projeto iremos ter necessidades em diversas dimensões e que teremos que utilizar diversas libs de terceiros ou construir as nossas próprias do zero.


O que é SHA-1

SHA-1 produz um valor de dispersão de 160 bits (20 bytes) conhecido como resumo da mensagem. Um valor de dispersão SHA-1 é normalmente tratado como um número hexadecimal de 40 dígitos.
Conhecido também como: "algoritmo de dispersão seguro" 


O que é CRC32

A verificação cíclica de redundância (do inglês, CRC - Cyclic Redundancy Check) é um método de detecção de erros normalmente usada em redes digitais e dispositivos de armazenamento

