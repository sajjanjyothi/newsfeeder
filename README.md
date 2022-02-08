# News Feeder
Service for providing news on various categories such as uk/technology from external sources

# How to run 
Service can be launched with docker compose as

docker-compose up --build

# Supported APIs
API's can be found in api/news.yaml in Open API spec 3 version, Which can be viewed on any API viewer 

# Authentication scheme for API
At the moment it supports bearer authentication which is exposed as environment variable in docker-compose , This can be changed to vault or secrets within kubernetes cluster


# API Examples

To get technology news 

```
curl -X GET http://127.0.0.1:30005/api/newsfeeder/v1/news/technology   -H "Authorization: Bearer ziglu" -H 'accept: application/json' 
```
To get all configured urls for categories
```
curl -X GET http://127.0.0.1:30005/api/newsfeeder/v1/news/urls   -H "Authorization: Bearer ziglu" -H 'accept: application/json'
```

# Helm charts for deployment
Can be improved and changed according to cluster configuration

```
make deploy 
```
