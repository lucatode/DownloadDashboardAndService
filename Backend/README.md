# Go Microservice With MongoDB

### Require MongoDB run locally
#### //TODO: make it configurable

## GET /downloads

``` bash
    [downloads:{_id,appid,latitude,longitude,country}, ...]
```

## POST /downloads

``` bash
    downloads:{_id,appid,latitude,longitude}
```

## GET /countDownloadsByCountry

``` bash
    [downloadsCount:{country,Count}, ...]
```

## GET /downloadsByCountry/:countryCode

``` bash
    [downloadsCount:{country,Count}]
```

### //TODO: GET /downloadsByTime/:dayTime

### //TODO: GET /countDownloadsByTime
