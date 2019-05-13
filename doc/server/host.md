```
curl -v 'http://localhost:9000/host/create' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "groupName": "test",
      "name": "test01",
      "ip": "192.168.1.1",
      "processor": "xeon",
      "memory": 4000000,
      "disk": 120000000,
      "platform": "Ubuntu 14.04"
    }'

curl -v -G 'http://localhost:9000/host/get' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1' \
  --data-urlencode 'name=test01' \
  --data-urlencode 'ip=192.168.1.1'

curl -v -G 'http://localhost:9000/host/list_all' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'order=name' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'groupName=test'

curl -v -G 'http://localhost:9000/host/list' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'page=0' \
  --data-urlencode 'size=2' \
  --data-urlencode 'order=name' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'groupName=test'

curl -v -G 'http://localhost:9000/host/count' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'groupName=test'

curl -v -G 'http://localhost:9000/host/exist' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1' \
  --data-urlencode 'name=test' \
  --data-urlencode 'ip=192.168.1.1'

curl -v 'http://localhost:9000/host/update' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1,
      "groupName": "test",
      "name": "test01",
      "ip": "192.168.1.1",
      "processor": "xeon x2",
      "memory": 12000000,
      "disk": 25600000,
      "platform": "Ubuntu 14.04.3 LTS"
    }'

curl -v 'http://localhost:9000/host/delete' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'
```