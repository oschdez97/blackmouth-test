### Blackmouth test exercise

---
author: [Oscar Hernandez](https://www.linkedin.com/in/oscar-luis-hernández-solano-905666225)
[github](https://github.com/oschdez97)
---

Run the API services
```shell
docker-compose up --build
```

Run tests
```shell
docker compose exec api bash
cd test && go test
```

API Documentation 
```
localhost:8080/docs/index.html
```

Shut down the API services
```shell
docker-compose down -v
```


