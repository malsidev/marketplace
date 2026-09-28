from aiokafka import AIOKafkaProducer

class KafkaProducer:
    def __init__(self, bootstrap_server: str):
       self.bootstrap_server = bootstrap_server
       self.producer : AIOKafkaProducer| None = None 

    async def start(self):
        self.producer = AIOKafkaProducer(
            bootstrap_servers=self.bootstrap_server
        )
        await self.producer.start()

    async def stop(self):
        if self.producer:
            await self.producer.stop()

    async def send(self, topic: str, message: dict):
        
        if not self.producer:
            raise RuntimeError("kafka producer is not started")
        await self.producer.send_and_wait(
            topic,
            message) 

producer = KafkaProducer("localhost:9092")