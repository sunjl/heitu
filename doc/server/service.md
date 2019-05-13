```
curl -v 'http://localhost:9000/service/create' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "groupName": "company",
      "projectId": 1,
      "name": "test",
      "protocol": "http",
      "hostName": "test01",
      "ip": "192.168.1.1",
      "port": 9000
    }'

curl -v -G 'http://localhost:9000/service/get' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1'

curl -v -G 'http://localhost:9000/service/list_all' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'order=projectId' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'groupName=company' \
  --data-urlencode 'projectId=1' \
  --data-urlencode 'projectName=test-server' \
  --data-urlencode 'protocol=http' \
  --data-urlencode 'hostName=test01' \
  --data-urlencode 'ip=192.168.1.1' \
  --data-urlencode 'port=9000'

curl -v -G 'http://localhost:9000/service/list' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'page=0' \
  --data-urlencode 'size=2' \
  --data-urlencode 'order=projectId' \
  --data-urlencode 'direction=asc' \
  --data-urlencode 'groupName=company' \
  --data-urlencode 'projectId=1' \
  --data-urlencode 'protocol=http' \
  --data-urlencode 'hostName=test01' \
  --data-urlencode 'ip=192.168.1.1' \
  --data-urlencode 'port=9000'

curl -v -G 'http://localhost:9000/service/count' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'groupName=company' \
  --data-urlencode 'projectId=1' \
  --data-urlencode 'protocol=http' \
  --data-urlencode 'hostName=test01' \
  --data-urlencode 'ip=192.168.1.1' \
  --data-urlencode 'port=9000'

curl -v -G 'http://localhost:9000/service/exist' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1'

curl -v 'http://localhost:9000/service/update' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1,
      "groupName": "company",
      "projectId": 1,
      "name": "test",
      "protocol": "http",
      "hostName": "test01",
      "ip": "192.168.1.1",
      "port": 9001
    }'

curl -v 'http://localhost:9000/service/delete' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'
```