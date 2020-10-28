# Example Go with Docker

### Intruksi

#### Build Image Manually
```bash
 docker build -t pasarkode/api:versi.1.1 .
```

#### Build using docker compose
``` bash
version: "3.3"
services:
  web:
    build: ./
    environment:
      - GIN_MODE=debug
    ports:
      - "8180:8180"

```

compose up
``` bash
docker-compose up
```
