import pika

# Use the 'rabbitmq-server' reference to point to the RabbitMQ container.
# If you're not using Docker Compose just put in the IP of the RabbitMQ broker instead.
connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq-server'))
channel = connection.channel()

channel.queue_declare(queue='hello', durable=True)

channel.basic_publish(exchange='',
                      routing_key='hello',
                      body='Hello World!',
                      properties=pika.BasicProperties(
                         delivery_mode = 2, # make message persistent
                      ))
print(" [x] Sent 'Hello World!'")

connection.close()
