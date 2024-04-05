# Prova 2 - Módulo 9

## Enunciado

- Implementar um Producer (Produtor): Deve coletar dados simulados de sensores de qualidade do ar e publicá-los em um tópico do Kafka chamado qualidadeAr. Os dados devem incluir:
  - Id do sensor, timestamp, tipo de poluente e nivel da medida.
- Implementar um Consumer (Consumidor): Deve assinar o tópico qualidadeAr e processar os dados recebidos, exibindo-os em um formato legível, além de armazená-los para análise posterior (escolha a forma de armazenamento que achar mais adequada).
- Implementar testes de Integridade e Persistência: Criar testes automatizados que validem a integridade dos dados transmitidos (verificando se os dados recebidos são iguais aos enviados) e a persistência dos dados (assegurando que os dados continuem acessíveis para consulta futura, mesmo após terem sido consumidos).

# Resolução

Esse repositório contém as seguintes pastas:

- `assests`: Contém os arquivos auxiliares, como um json com os dados de um sensor.
- `build`: Contém os arquivos de configuração do ambiente de execução (nesse caso apenas um docker-compose)
- `cmd`: Contém os entrypoints de cada pacote principal (nesse caso temos o _consumer_ e o _producer_)
- `internal`: Contém os pacotes internos dessa solução (e.g.: pacote de conexão com o mongo, kafka, etc)

## Execução

> [!IMPORTANT]  
> Para executar essa solução é necessário ter o docker, docker-compose e _golang_ instalados na máquina. _*Caso deseje rodar localmente, rode os comandos abaixo na root desse projeto!*_

### Como rodar

Primeiro, clone esse repositório e entre na pasta do projeto:

```bash
git clone https://github.com/Lemos1347/inteli-modulo-9-prova-2.git ; cd inteli-modulo-9-prova-2
```

Vamos subir o um cluster kafka e um mongodb localmente (para isso entre na pasta `build`):

```bash
cd build ; docker-compose up --build
```

Agora, vamos rodar o _producer_ (execute na root desse projeto):

```bash
go run cmd/producer/main.go
```

E por fim, vamos rodar o _consumer_ (execute na root desse projeto):

```bash
go run cmd/consumer/main.go
```

Pronto! Agora no terminal onde você rodou o _consumer_ você verá os dados sendo consumidos e salvos no banco de dados, e no terminal onde você rodou o _producer_ você verá os dados sendo produzidos e enviados ao kafka.

## Testes

Foram criados 2 testes para esse projeto:

- Teste de integridade - garante que os dados enviados pelo producer chegam ao tópico do Kafka sem alterações.
- Teste de persistência - garante que os dados enviados se mantem disponíveis para acesso em disco mesmo após o seu consumo inicial.

Ambos podem ser localizados no arquivo [`internal/kafka/kafka_test.go`](internal/kafka/kafka_test.go).

Para o teste de integridade, foi publicado no topico `prova_test` uma mensagem que é bytes de um json de um dado de um sensor e logo em seguida é verificado se os dados em memoria são iguais aos dados recebidos do kafka.  
Para o teste de persistência, foi publicado no tópico `prova_test_presistence` uma mensagem e posteriormente 2 consumers com group_id diferentes foram criados, possibilitando consumir a mesma mensagem 2 vezes. Após isso, foi verificado se a mensagem recebida em ambos os consumers eram iguais.
