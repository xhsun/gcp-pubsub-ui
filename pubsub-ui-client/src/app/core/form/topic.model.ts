import { FormControl, FormGroup, Validators } from "@angular/forms";
import { InputBase } from "./input-base.model";

export class Topic {
  projectID: InputBase<string>;
  topicName: InputBase<string>;

  constructor(projectID?: string, topicName?: string) {
    this.projectID = new InputBase<string>({
      value: projectID,
      key: 'projectID',
      label: 'GCP Project ID',
      validators: Validators.required
    });

    this.topicName = new InputBase<string>({
      value: topicName,
      key: 'topicName',
      label: 'GCP PubSub Topic',
      validators: Validators.required
    });
  }

  /**
   * Convert this topic model to a form group
   * @returns FormGroup
   */
  toFormGroup() {
    const group: any = {};
    group[this.projectID.key] = new FormControl(this.projectID.value || '', this.projectID.validators);
    group[this.topicName.key] = new FormControl(this.topicName.value || '', this.topicName.validators)
    return new FormGroup(group);
  }

  /**
   * Extract topic information from the provided form group
   * @param form FormGroup that contains topic information
   */
  toTopic(form: FormGroup) {
    if (!!form) {
      this.projectID.value = form.get(this.projectID.key)?.value;
      this.topicName.value = form.get(this.topicName.key)?.value;
    }
  }
}
