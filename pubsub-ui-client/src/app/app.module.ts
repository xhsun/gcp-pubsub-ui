import {  NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ReactiveFormsModule } from '@angular/forms';
import { PubSubUIClient } from './core/pubsubui/Pubsub_uiServiceClientPb';
import { TopicComponent } from './topic/topic.component';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatIconModule} from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';

const rpcClientFactory = () => new PubSubUIClient('http://localhost:9090');

@NgModule({
  declarations: [AppComponent, TopicComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatExpansionModule,
    MatIconModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule
  ],
  providers: [{ provide: PubSubUIClient, useFactory: rpcClientFactory }],
  bootstrap: [AppComponent],
})
export class AppModule {}
