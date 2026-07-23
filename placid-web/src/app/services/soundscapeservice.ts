import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';

@Service()
export class SoundscapeService {
  private http = inject(HttpClient);
}
