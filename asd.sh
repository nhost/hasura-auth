URL="http://localhost:4000"
email="asdasdasd@asdasd.com"
password="qweqweqweasasdasd"
uid="eedc3380-3b7b-47b8-a11d-a8fa69c2d26d"

# curl \
#     -H "Content-Type: application/json" \
#     -d '{"email": "'$email'", "password": "'$password'"}' \
#     -X POST $URL/signup/email-password

output=$(curl \
    -H "Content-Type: application/json" \
    -d '{"email": "'$email'", "password": "'$password'"}' \
    -X POST $URL/signin/email-password)

token=$(echo $output | jq -r '.session.accessToken')

# Get Orgs
curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $token" \
    -X GET $URL/organizations

# Create ORG
output=$(curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $token" \
    -d '{"name": "My Org"}' \
    -X POST $URL/organizations)

org_id=$(echo $output | jq -r '.id')

curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $token" \
    -d '{"name": "My Org"}' \
    -X POST $URL/organizations


# Get Orgs
curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $token" \
    -X GET $URL/organizations

curl \
    -H "Content-Type: application/json" \
    -d '{"email": "'$email'", "password": "'$password'"}' \
    -X POST $URL/signin/email-password

# delete
curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $token" \
    -X DELETE $URL/organizations/$org_id

curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $token" \
    -X DELETE $URL/organizations/$org_id

