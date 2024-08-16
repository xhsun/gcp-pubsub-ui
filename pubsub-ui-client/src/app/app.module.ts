import {  InjectionToken, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ReactiveFormsModule } from '@angular/forms';
import { TopicComponent } from './topic/topic.component';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatIconModule} from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import { environment } from 'src/environments/environment';
import { createConnectTransport } from '@connectrpc/connect-web';
import { createPromiseClient, PromiseClient } from '@connectrpc/connect';
import { PubSubUI } from './core/pubsubui/pubsub_ui_connect';

export const RPC_CLIENT = new InjectionToken<PromiseClient<typeof PubSubUI>>('RPC client');
const rpcClientFactory = () => {
  const transport = createConnectTransport({
    baseUrl: environment.apiUrl,
  });
  return createPromiseClient(PubSubUI, transport)};

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
    MatCardModule,
    MatDividerModule
  ],
  providers: [{ provide: RPC_CLIENT, useFactory: rpcClientFactory }],
  bootstrap: [AppComponent],
})
export class AppModule {}
