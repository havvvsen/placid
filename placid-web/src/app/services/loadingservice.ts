import { signal } from '@angular/core';
import { Service } from '@angular/core';

@Service({
  autoProvided: true,
})
export class LoadingService {
  isLoading = signal(false);

  show() {
    this.isLoading.set(true);
  }

  hide() {
    this.isLoading.set(false);
  }
}
