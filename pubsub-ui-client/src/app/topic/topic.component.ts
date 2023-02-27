import { Component, EventEmitter, Input, OnDestroy, OnInit, Output } from '@angular/core';
import { ClientReadableStream } from 'grpc-web';
import { PubSubUIClient } from '../core/pubsubui/Pubsub_uiServiceClientPb';
import * as rpcType from '../core/pubsubui/pubsub_ui_pb';

@Component({
  selector: 'app-topic',
  templateUrl: './topic.component.html',
  styleUrls: ['./topic.component.scss']
})
export class TopicComponent implements OnInit {
  decoder: TextDecoder = new TextDecoder();
  @Input()
  topicIndex!: number;
  @Input()
  projectID!: string;
  @Input()
  topicName!: string;
  @Output()
  shouldRemove = new EventEmitter<number>();

  stream!: ClientReadableStream<rpcType.Message>
  isEnd: boolean=false;
  lastUpdate: Date | undefined;
  messages: string[]=[]

  constructor(private rpcClient: PubSubUIClient){}

  ngOnInit() {
    const request = new rpcType.TopicSubscription().setGcpProjectId(this.projectID).setPubsubTopicName(this.topicName);
    this.stream = this.rpcClient.fetch(request);
    this.stream.on('status', function(status) {
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });
    this.stream.on("data", (message)=>{
      const timestamp = message.getTimestamp()?.toDate();
      const data = message.getData_asU8()
      this.lastUpdate = timestamp || this.lastUpdate;
      if (data.length > 0){
        this.isEnd = false;
        this.messages.push(`${timestamp?.toLocaleTimeString()} ${this.decoder.decode(data)}`)
      }
    });
    this. stream.on("error", (err)=>{
      this.isEnd = true;
      this.messages.push(err.message);
    });
    this.stream.on("end", ()=>{
      this.isEnd = true;
      this.messages.push("Connection closed, no more data to display");
    });
  }

  ngOnDestroy(): void {
    this.stream.cancel();
  }

  removeTopic(){
    this.stream.cancel();
    this.shouldRemove.emit(this.topicIndex);
  }
}
