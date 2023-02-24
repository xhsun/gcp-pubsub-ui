import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class TopicSubscription extends jspb.Message {
  getGcpProjectId(): string;
  setGcpProjectId(value: string): TopicSubscription;

  getPubsubTopicName(): string;
  setPubsubTopicName(value: string): TopicSubscription;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopicSubscription.AsObject;
  static toObject(includeInstance: boolean, msg: TopicSubscription): TopicSubscription.AsObject;
  static serializeBinaryToWriter(message: TopicSubscription, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopicSubscription;
  static deserializeBinaryFromReader(message: TopicSubscription, reader: jspb.BinaryReader): TopicSubscription;
}

export namespace TopicSubscription {
  export type AsObject = {
    gcpProjectId: string,
    pubsubTopicName: string,
  }
}

export class Message extends jspb.Message {
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): Message;

  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): Message;
  hasTimestamp(): boolean;
  clearTimestamp(): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    data: Uint8Array | string,
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

