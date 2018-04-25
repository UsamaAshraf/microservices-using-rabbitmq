import pika
import json

def emit_user_profile_update(user_id, new_data):
    # 'rabbitmq-server' is the network reference we have to the broker, 
    # thanks to Docker Compose.
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='rabbitmq-server'))
    channel    = connection.channel()

    exchange_name = 'user_updates'
    routing_key   = 'user.profile.update'

    # This will create the exchange if it doesn't already exist.
    channel.exchange_declare(exchange=exchange_name, exchange_type='topic', durable=True)
    
    new_data['id'] = user_id

    channel.basic_publish(exchange=exchange_name,
                          routing_key=routing_key,
                          body=json.dumps(new_data),
                          # Delivery mode 2 makes the broker save the message to disk.
                          properties=pika.BasicProperties(
                            delivery_mode = 2,
                        ))

    print("%r sent to exchange %r with data: %r" % (routing_key, exchange_name, new_data))
    connection.close()
