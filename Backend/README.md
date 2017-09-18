# Go Microservice With MongoDB

#### using MongoDB on mLab

### GET /downloads - https://gomicroservice.herokuapp.com/downloads

``` bash
    [downloads:{_id,appid,latitude,longitude,country}, ...]
```

### POST /downloads
#### Note: on insert coordinates, the service try to get country

``` bash
    downloads:{_id,appid,latitude,longitude}
```

### GET /countDownloadsByCountry - https://gomicroservice.herokuapp.com/downloads

``` bash
    [downloadsCount:{country,Count}, ...]
```

### GET /countDownloadsByCountryDetails - https://gomicroservice.herokuapp.com/downloadsByCountryDetail

``` bash
    [downloadsCount:{CountryDetail{...},Count}]
```


### GET /vueTableData - https://gomicroservice.herokuapp.com/vueTableData


### //TODO: GET /downloadsByTime/:dayTime

### //TODO: GET /countDownloadsByTime
