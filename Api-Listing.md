# cloud-watch-metric api-listing 

This markdown file contains all api document Order-wise how does flow works of cloud-watch-metric

	baseMetricUrl:
		http://localhost:7008

# API Reference
The cloud-watch-metric is organized around REST. Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.

## Errors

Cloud watch metric uses conventional HTTP response codes to indicate the success or failure of an API request. In general: Codes in the 2xx range indicate success. Codes in the 4xx range indicate an error that failed given the information provided (e.g., a required parameter was omitted, a charge failed, etc.). Codes in the 5xx range indicate an error with Stripe's servers (these are rare).

Some 4xx errors that could be handled programmatically (e.g., a card is declined) include an error code that briefly explains the error reported.

 # HTTPS STATUS CODE SUMMRY

Code   | Summary
------------- | -------------
200 - OK  | Everything worked as expected.
400 - Bad Request  | The request was unacceptable, often due to missing a required parameter.
401 - Unauthorized | No valid API key provided.
402 - Request Failed | The parameters were valid but the request failed.
403 - Forbidden | The API key doesn't have permissions to perform the request.
404 - Not Found | The requested resource doesn't exist.
409 - Conflict | The request conflicts with another request (perhaps due to using the same idempotent key)
429 - Too Many Requests | Too many requests hit the API too quickly. We recommend an exponential backoff of your requests.
500, 502, 503, 504 - Server Errors | Something went wrong on Stripe's end. (These are rare.)



**/awsx-metric/metric**
1. Api to post metric data using cloud element by id (get single query).

	Method: POST
	  Body:
		  request* String cloudElementId
      request* String cloudElementApiUrl
      request* String cloudWatchQueries
      
	Response:
		Success	Metric get records successfully response

*CURL*
```
curl --location 'http://localhost:7008/awsx-metric/metric/multi-queries' \
--header 'Content-Type: application/json' \
--data '{
   "cloudElementId":"",
   "cloudElementApiUrl":"",
    "cloudWatchQueries": "[{\"RefID\": \"A\",\"MaxDataPoint\": 100,\"Interval\": 60,\"TimeRange\": {\"From\": \"\",\"To\": \"\",\"TimeZone\": \"UTC\"},\"Query\": [{\"Namespace\": \"AWS/EC2\",\"MetricName\": \"CPUUtilization\",\"Period\": 300,\"Stat\": \"Average\",\"Dimensions\": [{\"Name\": \"InstanceId\",\"Value\": \"i-05e4e6757f13da657\"}]}]}]"
}'
```

#### OR

 Api to post metric data using zone,accessKey,secretKey,externalId,crossAccountRoleArn

Method: POST
	  Body:
		  request* String zone
      request* String accessKey
      request* String secretKey
      request* String externalId
      request* String crossAccountRoleArn
      request* String cloudWatchQueries
	Response:
		Success	Metric get records successfully response

*CURL*
```
curl --location 'http://localhost:7008/awsx-metric/metric/multi-queries' \
--header 'Content-Type: application/json' \
--data '{
    "zone": "us-east-1",
    "accessKey": "your accessKey",
    "secretKey": "your secretKey",
    "externalId": "your externalId",
    "crossAccountRoleArn": "your crossAccountRoleArn",
    "cloudWatchQueries": "[{\"RefID\": \"A\",\"MaxDataPoint\": 100,\"Interval\": 60,\"TimeRange\": {\"From\": \"\",\"To\": \"\",\"TimeZone\": \"UTC\"},\"Query\": [{\"Namespace\": \"AWS/EC2\",\"MetricName\": \"CPUUtilization\",\"Period\": 300,\"Stat\": \"Average\",\"Dimensions\": [{\"Name\": \"InstanceId\",\"Value\": \"i-05e4e6757f13da657\"}]}]}]"
}'
```



**/awsx-metric/metric/multi-queries**

2. Api to post metric data using cloud element by id (get multiple query).

	Method: POST
	  Body:
		  request* String cloudElementId
      request* String cloudElementApiUrl
      request* String cloudWatchQueries
      
	Response:
		Success	Metric get records successfully response

*CURL*
```
curl --location 'http://localhost:7008/awsx-metric/metric/multi-queries' \
--header 'Content-Type: application/json' \
--data '{
   "cloudElementId":"",
   "cloudElementApiUrl":"",
    "cloudWatchQueries": "[{\"RefID\": \"A\",\"MaxDataPoint\": 100,\"Interval\": 60,\"TimeRange\": {\"From\": \"\",\"To\": \"\",\"TimeZone\": \"UTC\"},\"Query\": [{\"Namespace\": \"AWS/EC2\",\"MetricName\": \"CPUUtilization\",\"Period\": 300,\"Stat\": \"Average\",\"Dimensions\": [{\"Name\": \"InstanceId\",\"Value\": \"i-05e4e6757f13da657\"}]}]}]"
}'
```

#### OR

 Api to post metric data using zone,accessKey,secretKey,externalId,crossAccountRoleArn

Method: POST
	  Body:
		  request* String zone
      request* String accessKey
      request* String secretKey
      request* String externalId
      request* String crossAccountRoleArn
      request* String cloudWatchQueries
	Response:
		Success	Metric get records successfully response

*CURL*
```
curl --location 'http://localhost:7008/awsx-metric/metric/multi-queries' \
--header 'Content-Type: application/json' \
--data '{
    "zone": "us-east-1",
    "accessKey": "your accessKey",
    "secretKey": "your secretKey",
    "externalId": "your externalId",
    "crossAccountRoleArn": "your crossAccountRoleArn",
    "cloudWatchQueries": "[{\"RefID\": \"A\",\"MaxDataPoint\": 100,\"Interval\": 60,\"TimeRange\": {\"From\": \"\",\"To\": \"\",\"TimeZone\": \"UTC\"},\"Query\": [{\"Namespace\": \"AWS/EC2\",\"MetricName\": \"CPUUtilization\",\"Period\": 300,\"Stat\": \"Average\",\"Dimensions\": [{\"Name\": \"InstanceId\",\"Value\": \"i-05e4e6757f13da657\"}]}]}]"
}'
```







