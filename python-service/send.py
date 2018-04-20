import pika

# Use the 'rabbitmq-server' reference to point to the RabbitMQ container. 
connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq-server'))
channel = connection.channel()

