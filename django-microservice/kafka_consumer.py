from confluent_kafka import Consumer, KafkaError

c = Consumer({
    'bootstrap.servers': 'kafka:9092',
    'group.id': 'django-consumer-group',
    'auto.offset.reset': 'earliest'
})

c.subscribe(['example-topic'])

while True:
    msg = c.poll(1.0)
    if msg is None:
        continue
    if msg.error():
        if msg.error().code() == KafkaError._PARTITION_EOF:
            continue
        else:
            print(msg.error())
            break

    print(f'Received message: {msg.value().decode("utf-8")}')

c.close()
