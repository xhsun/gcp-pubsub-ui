import { ValidatorFn, AbstractControlOptions } from '@angular/forms';

/**
 * The base input model to represent html input
 */
/*istanbul ignore file*/
export class InputBase<T> {
  /**
   * Value of current input
   */
  value: T|undefined;

  /**
   * Key used to retrieve current input from its parent form
   */
  key: string;

  /**
   * Label used to display to user
   */
  label: string;

  /**
   * Validators to apply to current input
   */
  validators?: ValidatorFn | ValidatorFn[] | AbstractControlOptions | null;

  constructor(
    options: {
      value?: T;
      key?: string;
      label?: string;
      validators?: ValidatorFn | ValidatorFn[] | AbstractControlOptions | null;
    } = {}
  ) {
    this.value = options.value;
    this.key = options.key || '';
    this.label = options.label || '';
    this.validators = options.validators;
  }
}
