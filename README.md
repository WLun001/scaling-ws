## Get started

- For initial problem, checkout `problem` branch
- For nats solution, checkout `solution-nats-ws` branch


## How to run
### `problem` branch
```bash
# on terminal
# start web ui
cd web
npm i && npm run serve

# open another terminal
# start ws server
go run main.go --addr :4000 --skipWs

# open another terminal
# start api server
go run main.go
```
Open two tabs on `http://localhost:8080`

### `solution-nats-ws` branch
```bash
# on terminal
# start web ui
cd web
npm i && npm run serve

# open another terminal
# start nats-server
# make sure it is installed
nats-server

# open another terminal
# start api server
make start-api

# open another terminal
# start ws server
make start-ws
```
Open two tabs on `http://localhost:8080`



