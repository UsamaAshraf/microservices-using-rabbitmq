import pika
import json

def emit_user_profile_update(user_id, new_data):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='rabbitmq-server'))
    channel = connection.channel()

    channel.exchange_declare(exchange='user_updates', exchange_type='topic')

    channel.basic_publish(exchange='topic_logs',
                          routing_key='user.profile.update',
                          body=json.dumps(new_data))

    print("Sent %r:%r:%r" % ('topic_logs', 'user.profile.update', new_data))
    connection.close()
