# Organizando seu projeto Go.


Vamos criar pacotes Go, vamos tentar entender um pouco quando estamos criando projetos em Go.


Tudo inicia em Go mod, o que temos que saber e como usa-lo.

 - go mod init 
 	- Cria um novo módulo, inicializando o go.mod

 - go build
 	- Testa outros comandos de construção de pacote adicionam novas dependências go.mod conforme necessário.

 - go list -m all 
 	- Imprime as dependências do módulo atual.

 - go get 
 	- Altera a versão necessária de uma dependência (ou adiciona uma nova dependência).

 - go mod tidy 
 	- Remove dependências não utilizadas.

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


### Go mod


#### GOPROXY
	
	Essa variável permite definir um Go module proxy, que basicamente vai ser uma aplicação que vai intermediar sua requisição de dependências. 
	O valor padrão é https://proxy.golang.org mas você pode fazer uso de outros, como o JFrog por exemplo ou colocar off

	Um GOPROXY controla a origem dos downloads do seu módulo Go e pode ajudar a garantir que as compilações sejam determinísticas e seguras.

	Definir um GOPROXY para o desenvolvimento Golang ou ambiente CI redireciona as solicitações de download do módulo Go para um repositório de cache. 

	Um GOPROXY público é um repositório centralizado disponível para desenvolvedores Golang em todo o mundo. Ele hospeda módulos Go de código aberto que foram disponibilizados por terceiros em repositórios de projetos.

	Além de realizar downloads, um GOPROXY público também pode fornecer aos desenvolvedores GoLang informações mais detalhadas sobre os módulos que contém. O JFrog GoCenter oferece uma interface do usuário rica com a capacidade de pesquisar e acessar informações de segurança, como CVEs , metadados não relacionados à segurança, como estatísticas de adoção e suporte gosumdb. Esses metadados ajudam os usuários a tomar melhores decisões ao selecionar um módulo Go público. 

	Este uso de GOPRIVATE também garante que o uso desses módulos privados não "vaze" por meio de solicitações para um servidor de banco de dados público GOPROXY & checksum em uma rede aberta. Outra alternativa é usar a variável GONOSUMDB que inclui referências a módulos go privados. Embora essa configuração permita que o cliente Go resolva dependências de módulo público e privado, ela não impõe requisitos de imutabilidade ou disponibilidade para módulos privados.

#### GOPROXY privado

Um GOPROXY privado é aquele que você instala para armazenar módulos Go públicos e privados em sua própria infraestrutura.

Módulos públicos são armazenados em cache localmente por proxy de um GOPROXY público em um gerenciador de repositório binário como JFrog Artifactory. Módulos privados também são armazenados em cache em um repositório de seus repositórios VCS. Dessa forma, a imutabilidade e a disponibilidade podem ser garantidas para os módulos Go públicos e privados. 
	
	Exemplo:
```bash
	$ export GOPROXY="https://:@my.artifactory.server/artifactory/api/go/go
	$ export GONOSUMDB="github.com/mycompany/*, github.com/mypersonal/*"
```

#### GONOSUMDB

Essa variável serve para você dizer que as suas dependências não devem ser verificadas pelo checksum.


#### GOSUMDB

Quando você faz o download de uma dependência, como você sabe que ela é verdadeira? Como garantir que ninguém modificou ela inserindo algum bug (ou código malicioso) nela?
Uma prática comum de mercado é fazer uso do checksum para isso, onde basicamente nós fazemos uma verificação em uma fonte confiável se a integridade da nossa dependência é válida. O Golang faz isso por padrão, onde o valor dessa env por padrão é sum.golang.org, um database gratuito oferecido pelo google.


#### GONOPROXY

Essa variável serve para você dizer que as dependências devem ser buscadas diretamente do controle de versão (Github em muitos casos, mas pode ser qualquer um) sem passar pelo Proxy. Suas dependências privadas por exemplo, não vão estar no repositório público, então elas vão ser recuperadas diretamente do controle de versão.

#### GONOSUMDB

Essa variável serve para você dizer que as suas dependências não devem ser verificadas pelo checksum. Muito cuidado ao usar isso, pois aqui pode haver uma brecha de segurança. Você pode fazer uso das suas dependências privadas aqui, mas confiar em terceiros pode facilmente se tornar um problema crítico de segurança.


#### GOPRIVATE

Se você estava prestando atenção, percebeu que para fazer uso de dependências privadas precisaria definir a variável GONOPROXY e GONOSUMDB, evitando assim o uso do proxy e do checksum. Isso não é necessário, definindo apenas essa variável você obtém essa mesmo comportamento, pois quando elas não tiverem nenhum valor definido elas vão assumir por padrão o valor definido em GOPRIVATE.


O que é SHA-1

SHA-1 produz um valor de dispersão de 160 bits (20 bytes) conhecido como resumo da mensagem. Um valor de dispersão SHA-1 é normalmente tratado como um número hexadecimal de 40 dígitos.
Conhecido também como: "algoritmo de dispersão seguro" 


O que é CRC32

A verificação cíclica de redundância (do inglês, CRC - Cyclic Redundancy Check) é um método de detecção de erros normalmente usada em redes digitais e dispositivos de armazenamento

#### GOINSECURE

Por padrão a conexão de módulos precisa fazer uso de conexões HTTPS e validar o certificado, o que é ótimo para segurança, mas pode te atrapalhar em desenvolvimento ou até ao fazer uso de alguma PoC.
Essa variável permite você definir valores de servidores que confia para pular a exigência de HTTPS e de certificados válidos. Aqui vale o mesmo cuidado já citado no caso do GONOSUMDB, cuidado ao usar. Em produção, jamais.
