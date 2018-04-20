import pika

# Use the 'rabbitmq-server' reference to point to the RabbitMQ container.
# If you're not using Docker Compose just put in the IP of the RabbitMQ broker instead.
connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq-server'))
channel = connection.channel()

channel.queue_declare(queue='hello', durable=True)

def callback(ch, method, properties, body):
    print("Received %r" % ch)
    print("Received %r" % method)
    print("Received %r" % properties)
    print("Received %r" % body)
    # Not needed when no_ack=True
    ch.basic_ack(delivery_tag = method.delivery_tag)


#channel.basic_consume(callback, queue='hello', no_ack=True)
channel.basic_consume(callback, queue='hello')

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
