## Login


<details>
  <summary>Login Request</summary>
  
```json
    {
        "email": "example@gmail.com",
        "password": "example-password"
    }
```
</details>

<details>
    <summary>Login Response</summary>

```json
    {
        "status": "success",
        "message": "Login success",
        "data": {
            "token": "token jwt",
            "expired_at": "timestamp"
        }
    }
```
</details>
