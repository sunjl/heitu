```
curl -v 'http://localhost:9000/config/create' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "projectId": 1,
      "version": "1.0",
      "environment": "dev",
      "fileName": "dev.properties",
      "content": "key=value"
    }'
	
curl -v -G 'http://localhost:9000/config/get' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1'

curl -v -G 'http://localhost:9000/config/list_all' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'order=projectId' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'projectId=1' \
  --data-urlencode 'projectName=test-server' \
  --data-urlencode 'version=1.0' \
  --data-urlencode 'environment=dev'

curl -v -G 'http://localhost:9000/config/list' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'page=0' \
  --data-urlencode 'size=2' \
  --data-urlencode 'order=projectId' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'projectId=1' \
  --data-urlencode 'projectName=test-server' \
  --data-urlencode 'version=1.0' \
  --data-urlencode 'environment=dev'

curl -v -G 'http://localhost:9000/config/count' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'projectId=1' \
  --data-urlencode 'version=1.0' \
  --data-urlencode 'environment=dev'

curl -v -G 'http://localhost:9000/config/exist' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1'

curl -v 'http://localhost:9000/config/update' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1,
      "projectId": 1,
      "version": "1.0",
      "environment": "dev",
      "fileName": "dev.properties",
      "content": "key=value2"
    }'

curl -v 'http://localhost:9000/config/delete' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'
```