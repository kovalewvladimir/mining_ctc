# mining_ctc_bot
Telegram bot

# Utils
## install gb (vendoring )
```
go get github.com/constabulary/gb/...
```
### use gb
```
gb vendor fetch <importpath>
```

# Install
```
go get github.com/constabulary/gb/...
gb vendor restore
```

# Deploy

## heroku
Create app
```
heroku create mining-ctc-bot
```
Set variables app
```
heroku config:set test=123 --app mining-ctc-bot
```
Show variables app
```
heroku config --app mining-ctc-bot
```