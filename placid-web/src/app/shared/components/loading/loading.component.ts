import { LoadingService } from '@/services/loadingservice';
import { Component, inject } from '@angular/core';

@Component({
  selector: 'app-loading-component',
  templateUrl: 'loading.html',
})
export class LoadingComponent {
  loadingService = inject(LoadingService);
}
