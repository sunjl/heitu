```
curl -v 'http://localhost:9000/project/create' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "groupName": "test",
      "name": "test-server",
      "git": "git@git.example.com:test/test-server.git"
    }'

curl -v -G 'http://localhost:9000/project/get' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1' \
  --data-urlencode 'name=test-server'

curl -v -G 'http://localhost:9000/project/list_all' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'order=groupName' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'groupName=test'

curl -v -G 'http://localhost:9000/project/list' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'page=0' \
  --data-urlencode 'size=2' \
  --data-urlencode 'order=groupName' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'groupName=test'

curl -v -G 'http://localhost:9000/project/count' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'groupName=test'

curl -v -G 'http://localhost:9000/project/exist' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1' \
  --data-urlencode 'name=test-server'

curl -v 'http://localhost:9000/project/update' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1,
      "groupName": "test",
      "name": "test-server-new",
      "git": "git@git.example.com:test/test-server-new.git"
    }'

curl -v 'http://localhost:9000/project/delete' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'
```