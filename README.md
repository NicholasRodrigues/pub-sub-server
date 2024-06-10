# pub-sub-server

## Running RabbitMQ with Docker

To run RabbitMQ with Docker, use the following command:

```bash
  docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
```

The server itself is running on port 5672, but you won't interact with it directly through your browser, only through code. However, you can use the RabbitMQ management UI in your browser.

To access the management UI, open 

```bash
  http://localhost:15672.
```
