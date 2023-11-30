### `CLOUD-WATCH-AWSX-API Port: 7008`
### `1. get cloud watch metric (single query)`

```
Method - Post
```
```
API End Point - /awsx-metric/metric
Request -  Body {cloudElementId,cloudElementApiUrl,query} or {accessKey,secretKey,externalId,crossAccountRoleArn,cloudWatchQueries}
Response - GetMetricDataWithSingleQuery(w http.ResponseWriter, r *http.Request)
```

	1. Request passed to router  
	2. Router forward request to handlers
	3. Handlers pass to controller 
	4. Controller call function  &cloudwatch.MetricDataQuery 
	5. this function  return result, err
	6. If error is not nil then send response with status code and message
	NOTE--> result will be generated from a single query
<hr>

	





### `2. get cloud watch metric (multiple query)`

```
Method - Post
```
```
API End Point - /awsx-metric/metric/multi-queries
Request -  Body {cloudElementId,cloudElementApiUrl,query} or {accessKey,secretKey,externalId,crossAccountRoleArn,cloudWatchQueries}
Response - GetMetricDataWithMultipleQueries(w http.ResponseWriter, r *http.Request)
```

	1. Request passed to router  
	2. Router forward request to handlers
	3. Handlers pass to controller 
	4. Controller call function  &cloudwatch.MetricDataQuery 
	5. this function  return result, err
	6. If error is not nil then send response with status code and message
	NOTE--> result will be generated from a multiple query
<hr>



### Who do I talk to? ###

	Please mail us on
	info@syenctiks.com
Footer
© 2023 GitHub, Inc.
Footer navigation
Terms
Privacy
Security
Status
Docs
Contact GitHub
Pricing
API
Training
Blog
About






