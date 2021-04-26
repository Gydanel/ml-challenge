# ml-challenge

## Usage

For running the app locally execute in project root:
```
export AUTH_USER="user"
export AUTH_SECRET="secret"
make run
```

### Api
#### Topsecret
```
curl --location --request POST 'http://localhost:8080/ml-challenge/api/topsecret' \
     --U <user>:<secret>' \
     --header 'Content-Type: application/json' \
     --data-raw '{
          "satellites": [
              {
                  "name": "kenobi",
                  "distance": <float>,
                  "message": ["<string>", ...]
              },
                      {
                  "name": "sato",
                  "distance": <float>,
                  "message": ["<string>", ...]
              },
                      {
                  "name": "skywalker",
                  "distance": <float>,
                  "message": ["<string>", ...]
              }
          ]
     }'
```
#### Topsecret split
```
curl --location --request POST 'http://localhost:8080/ml-challenge/api/topsecret_split/<satellite_name>' \
     --U <user>:<secret>' \
     --header 'Content-Type: application/json' \
     --data-raw '{
        "distance": <float>,
        "message": ["<string>", ...]
     }'
```

```
curl --location --request GET 'http://localhost:8080/ml-challenge/api/topsecret_split' \
     --U <user>:<secret>' \
```
