```
curl -v 'http://localhost:9000/task/create' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "hostName": "test01",
      "ip": "192.168.1.1",
      "name": "任务名称",
      "content": "echo \"test\"",
      "status": "waiting"
    }'

curl -v -G 'http://localhost:9000/task/get' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1'

curl -v -G 'http://localhost:9000/task/list_all' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'order=hostName' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'hostName=test01' \
  --data-urlencode 'ip=192.168.1.1' \
  --data-urlencode 'status=waiting'

curl -v -G 'http://localhost:9000/task/list' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'page=0' \
  --data-urlencode 'size=2' \
  --data-urlencode 'order=hostName' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'hostName=test01' \
  --data-urlencode 'ip=192.168.1.1' \
  --data-urlencode 'status=waiting'

curl -v -G 'http://localhost:9000/task/count' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'hostName=test01' \
  --data-urlencode 'ip=192.168.1.1' \
  --data-urlencode 'status=waiting'

curl -v -G 'http://localhost:9000/task/exist' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1'

curl -v 'http://localhost:9000/task/update' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1,
      "hostName": "test01",
      "ip": "192.168.1.1",
      "name": "任务名称",
      "content": "echo \"test new\"",
      "status": "done"
    }'

curl -v 'http://localhost:9000/task/delete' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'
```