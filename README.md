# Welcome to todolist-go

A [Go](http://www.golang.org/)/[Revel](https://revel.github.io/manual/index.html) web application with an [AngularJS](https://angularjs.org/) frontend.


### Build and run the application:
```
docker run --name todolist -p 9000:9000 -it debian:latest
apt update && apt-get install -y wget git gcc npm
wget -qO- https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz | tar xzf - -C /opt/
export GOROOT=/opt/go ; export GOPATH=/usr/share/gocode ; export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
go get -v -d github.com/denisacostaq/todolist-go/...
go get -v -u github.com/revel/cmd/revel
cd $GOPATH/src/github.com/denisacostaq/todolist-go/public ; npm i
revel run github.com/denisacostaq/todolist-go
```

### Go to http://localhost:9000/public/app/index.html#!/tasks and you'll see:

    A tasks list (empty if you have not any yet).

### Creating some test data
```
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "tests", "priority": 10}' http://localhost:9000/api/v1.0/labels
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "design", "priority": 7}' http://localhost:9000/api/v1.0/labels
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "architecture", "priority": 5}' http://localhost:9000/api/v1.0/labels
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "i18n", "priority": 2}' http://localhost:9000/api/v1.0/labels
curl --request GET --header "Accept: application/json" http://localhost:9000/api/v1.0/labels

curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "task 1", "due_date": "2019-10-28T00:00:00.000Z", "labels": [{"id": 1}]}' http://localhost:9000/api/v1.0/tasks
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "task 2", "due_date": "2019-10-30T00:00:00.000Z", "labels": [{"id": 1}, {"id": 2}]}' http://localhost:9000/api/v1.0/tasks
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "task 3", "due_date": "2019-10-18T00:00:00.000Z", "labels": [{"id": 3}, {"id": 2}]}' http://localhost:9000/api/v1.0/tasks
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "task 4", "due_date": "2019-10-11T00:00:00.000Z", "labels": [{"id": 1}, {"id": 2}, {"id": 4}]}' http://localhost:9000/api/v1.0/tasks
curl --request POST --header "Accept: application/json" --header "Content-Type: application/json" --data '{"name": "task 5", "due_date": "2019-10-12T00:00:00.000Z", "labels": [{"id": 1}]}' http://localhost:9000/api/v1.0/tasks
curl --request GET --header "Accept: application/json" http://localhost:9000/api/v1.0/tasks
```

## Code Layout

The directory structure of the application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        business/     Business service definitions
        models/       Database model definition
        repository/   Data access layer to get data from satabase
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        app/          Angular app

    tests/            Test suites (not used yet)


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

