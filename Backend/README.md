# Go Microservice With MongoDB

#### using MongoDB on mLab

### GET /downloads - https://gomicroservices.herokuapp.com/downloads

``` bash
    [downloads:{_id,appid,latitude,longitude,country}, ...]
```

### POST /downloads
#### Note: on insert coordinates, the service try to get country

``` bash
    downloads:{_id,appid,latitude,longitude}
```

### GET /countDownloadsByCountry - https://gomicroservices.herokuapp.com/downloads

``` bash
    [downloadsCount:{country,Count}, ...]
```

### GET /countDownloadsByCountryDetails - https://gomicroservices.herokuapp.com/downloadsByCountryDetail

``` bash
    [downloadsCount:{CountryDetail{...},Count}]
```


### GET /vueTableData - https://gomicroservices.herokuapp.com/vueTableData


### //TODO: GET /downloadsByTime/:dayTime

### //TODO: GET /countDownloadsByTime
