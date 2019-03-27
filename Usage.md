# [Shredx](https://github.com/shredx) Redis Rate Limiter
APIs for Redis Rate Limiter platform built as part of https://github.com/shredx

The platform consists of three parts
* [Node Token Generator](https://github.com/shredx/node-redis-rate-limiter)
* [Golang Rate Limiter](https://github.com/shredx/golang-redis-rate-limiter)
* [Golang API Gateway](https://github.com/shredx/golang-api-gateway)

## Usage
* First create the user
* Then create the subscription key
* Then use the api to be limited
* View the usage status using API Usage
* Reset the token useage using Usage Reset

## Requests

### **POST** - /v1/users/

#### Description
This API will create a test user in the platform.
You have to now create a subscription key for that user

#### CURL

```sh
curl -X POST "http://127.0.0.1:9090/v1/users/" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -H "Cookie: REVEL_FLASH=" \
    --data-raw "email"="test@test.com" \
    --data-raw "name"="Tester"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/x-www-form-urlencoded"
  ],
  "default": "application/x-www-form-urlencoded"
}
```
- **Cookie** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "REVEL_FLASH="
  ],
  "default": "REVEL_FLASH="
}
```

#### Body Parameters

- **email** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "test@test.com"
  ],
  "default": "test@test.com"
}
```
- **name** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "Tester"
  ],
  "default": "Tester"
}
```

### **POST** - /v1/users/subscriptions

#### Description
This API will create a subscription key for a user with given email address. Now you can copy the token key and use it with the requests to the API gateway to hit the required api with header key for the token as `token`.

#### CURL

```sh
curl -X POST "http://127.0.0.1:9090/v1/users/subscriptions" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -H "Cookie: REVEL_FLASH=" \
    --data-raw "email"="test@test.com"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/x-www-form-urlencoded"
  ],
  "default": "application/x-www-form-urlencoded"
}
```
- **Cookie** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "REVEL_FLASH="
  ],
  "default": "REVEL_FLASH="
}
```

#### Body Parameters

- **email** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "test@test.com"
  ],
  "default": "test@test.com"
}
```

### **GET** - /

#### Description
You can check the usage of your API using the API Usage api

#### CURL

```sh
curl -X GET "http://127.0.0.1:8080/" \
    -H "token: e3795d31-6423-4f46-8acd-1724888955c5" \
    -H "Cookie: REVEL_FLASH="
```

#### Header Parameters

- **token** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "e3795d31-6423-4f46-8acd-1724888955c5"
  ],
  "default": "e3795d31-6423-4f46-8acd-1724888955c5"
}
```
- **Cookie** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "REVEL_FLASH="
  ],
  "default": "REVEL_FLASH="
}
```

### **POST** - /usage

#### Description
This api gives information about the your api usage.
You can reset the api usage using Usage Reset api

#### CURL

```sh
curl -X POST "http://127.0.0.1:8085/usage" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -H "Cookie: REVEL_FLASH=" \
    --data-raw "token"="e3795d31-6423-4f46-8acd-1724888955c5"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/x-www-form-urlencoded"
  ],
  "default": "application/x-www-form-urlencoded"
}
```
- **Cookie** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "REVEL_FLASH="
  ],
  "default": "REVEL_FLASH="
}
```

#### Body Parameters

- **token** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "e3795d31-6423-4f46-8acd-1724888955c5"
  ],
  "default": "e3795d31-6423-4f46-8acd-1724888955c5"
}
```

### **POST** - /reset

#### Description
This API will reset the usage of the given api token

#### CURL

```sh
curl -X POST "http://127.0.0.1:8085/reset" \
    -H "Content-Type: application/x-www-form-urlencoded; charset=utf-8" \
    -H "Cookie: REVEL_FLASH=" \
    --data-raw "token"="e3795d31-6423-4f46-8acd-1724888955c5"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/x-www-form-urlencoded; charset=utf-8"
  ],
  "default": "application/x-www-form-urlencoded; charset=utf-8"
}
```
- **Cookie** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "REVEL_FLASH="
  ],
  "default": "REVEL_FLASH="
}
```

#### Body Parameters

- **token** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "e3795d31-6423-4f46-8acd-1724888955c5"
  ],
  "default": "e3795d31-6423-4f46-8acd-1724888955c5"
}
```

## References

