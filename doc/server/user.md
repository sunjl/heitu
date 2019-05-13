```
curl -v 'http://localhost:9000/user/create' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "username": "test",
      "password": "test",
      "displayName": "显示名称",
      "phone": "123456",
      "userType": "admin"
    }'

curl -v http://localhost:9000/user/sign_in \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{
      "username": "test",
      "password": "test"
    }'

curl -v -G 'http://localhost:9000/user/get' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1' \
  --data-urlencode 'username=test' \
  --data-urlencode 'token=47bfa2b8-2468-4c6e-8712-a5a240ea6ea7'

curl -v -G 'http://localhost:9000/user/list_all' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'order=createTime' \
  --data-urlencode 'direction=desc' \
  --data-urlencode 'userType=admin'

curl -v -G 'http://localhost:9000/user/list' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'page=0' \
  --data-urlencode 'size=2' \
  --data-urlencode 'order=createTime' \
  --data-urlencode 'direction=desc' \
  --data-urlencode 'userType=admin'

curl -v -G 'http://localhost:9000/user/count' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
   --data-urlencode 'userType=normal'

curl -v -G 'http://localhost:9000/user/exist' \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --data-urlencode 'id=1' \
  --data-urlencode 'username=test' \
  --data-urlencode 'token=47bfa2b8-2468-4c6e-8712-a5a240ea6ea7'

curl -v 'http://localhost:9000/user/update' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1,
      "displayName": "显示名称new",
      "phone": "456789",
      "userType": "admin"
    }'

curl -v 'http://localhost:9000/user/update_password' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "username": "test",
      "newPassword": "new"
    }'

curl -v 'http://localhost:9000/user/update_token' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'

curl -v 'http://localhost:9000/user/delete' \
  --request POST \
  --header "X-Auth-Token: 47bfa2b8-2468-4c6e-8712-a5a240ea6ea7" \
  --header "Content-Type: application/json" \
  --data '{
      "id": 1
    }'
```