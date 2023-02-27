import { Component, ViewChild } from '@angular/core';
import { FormGroup, FormGroupDirective, NgForm } from '@angular/forms';
import { MatAccordion } from '@angular/material/expansion';
import { Topic } from './core/form/topic.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  @ViewChild(MatAccordion) accordion!: MatAccordion;
  @ViewChild('formDirective') private formDirective!: NgForm;

  topicKeyFormatter='{0}{1}';
  topics: Map<string, [string, string]> = new Map();
  currentTopic: Topic = new Topic();
  form!: FormGroup;


  ngOnInit() {
    this.form = this.currentTopic.toFormGroup();
  }

  onSubmit() {
    if (!!this.form && this.form.valid) {
      this.currentTopic.toTopic(this.form);
      const id = this.currentTopic.projectID.value || '';
      const topic = this.currentTopic.topicName.value || '';
      const key = this.stringFormat(this.topicKeyFormatter, id, topic);
      if (this.topics.has(key)) {
        this.form.controls[this.currentTopic.topicName.key].setErrors({ unique: true });
      } else {
        this.topics.set(key, [id, topic]);
        this.form.reset();
        this.formDirective.resetForm();
      }
    }
  }

  removeTopic(topic: [string, string]){
    this.topics.delete(this.stringFormat(this.topicKeyFormatter, topic[0], topic[1]));
  }

  stringFormat (str: string, ...args: string[]){
    return str.replace(/{(\d+)}/g, (_, index) => args[index] || '')
  }
}
