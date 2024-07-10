import { Injectable } from '@nestjs/common';
import { ClientKafka } from '@nestjs/microservices';
import { Inject } from '@nestjs/common';

@Injectable()
export class KafkaService {
  constructor(
    @Inject('KAFKA_CLIENT') private readonly kafkaClient: ClientKafka,
  ) {}

  async onModuleInit() {
    this.kafkaClient.subscribeToResponseOf('example-topic');
    await this.kafkaClient.connect();
  }

  async sendMessage(message: string) {
    this.kafkaClient.emit('example-topic', { message });
  }
}
