import { Component, EventEmitter, Inject, Input, OnInit, Output } from '@angular/core';
import { PromiseClient } from '@connectrpc/connect';
import { PubSubUI } from '../core/pubsubui/pubsub_ui_connect';
import { RPC_CLIENT } from '../app.module';
import { from, Subscription } from 'rxjs';

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

  stream!: Subscription;
  streamEnder = new AbortController();
  isEnd: boolean=false;
  lastUpdate: Date | undefined;
  messages: string[]=[]
  rpcClient: PromiseClient<typeof PubSubUI>

  constructor(@Inject(RPC_CLIENT) rpcClient: PromiseClient<typeof PubSubUI>){
    this.rpcClient = rpcClient;
  }

  ngOnInit() {
    this.stream = from(this.rpcClient.fetch({
      gcpProjectId: this.projectID,
      pubsubTopicName: this.topicName,
    }, {signal: this.streamEnder.signal})).subscribe({
      next:message => {
        const timestamp = message.timestamp?.toDate();
        const data = message.data;
        this.lastUpdate = timestamp || this.lastUpdate;
        if (data.length > 0) {
          this.isEnd = false;
          this.messages.push(
            `${timestamp?.toLocaleTimeString()} ${this.decoder.decode(data)}`
          );
        }
      },
      error: err => {
        this.isEnd = true;
        this.messages.push(err.message);
      },
      complete: () =>{
        this.isEnd = true;
        this.messages.push('Connection closed, no more data to display');
      }
    });
  }

  ngOnDestroy(): void {
    this.stream.unsubscribe();
    this.streamEnder.abort();
  }

  removeTopic(){
    this.shouldRemove.emit(this.topicIndex);
  }
}
