# How to run the code

## Run the server

```
make
```

## Make a request

```
curl http://localhost:8000/latest_volumes -H "Content-Type: application/json" -d '["B0C493GKLX","B09478443K","B074CFBX48"]'
```

# How to run the tests

```
make test
```

# TODO NEXT

[ ] Find a way to save data