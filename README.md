# typeperf-wsl-go

Making a bridge for using Windows hardware stats in WSL Linux

## POC

```shell
make poc

# or, if you have jq:

go run main.go | jq .
make poc | jq .
```
