# test-queroquitar

## Descrição
Webserver em Go que retorna uma página contendo um "greeter" e um gif animado randômico (gif do site giphy.com).

## Instalação
Para a instalação do projeto execute os comandos a seguir:

    mkdir $GOPATH/src/test/queroquitar

    git clone git@github.com:burredis/test-queroquitar.git

## Utilização
Acesse o diretório do projeto:

    cd $GOPATH/src/test/queroquitar

Iniciar o Docker:

    docker-compose up --build -d --remove-orphans --force-recreate

Acessar a URL via browser:

    http://localhost:8000/?name={VALUE}

  Querystring:
    * name: Nome para saudação

## Dependências
    https://echo.labstack.com/
    https://github.com/sanzaru/go-giphy