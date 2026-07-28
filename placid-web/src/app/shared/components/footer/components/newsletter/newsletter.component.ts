import { LoadingService } from '@/services/loadingservice';
import NewsletterService from '@/services/newsletterservice';
import { NotificationService } from '@/services/notificationservice';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'footer-newsletter-component',
  standalone: true,
  templateUrl: 'newsletter.html',
  imports: [FormsModule],
})
export class FooterNewsLetterComponent {
  http = inject(HttpClient);
  newsletterService = inject(NewsletterService);
  loadingService = inject(LoadingService);
  notificationService = inject(NotificationService);
  email = '';

  onSubmit() {
    this.loadingService.isLoading.set(true);

    if (!this.email.includes('@') || !this.email.includes('.') || this.email.length < 4) {
      this.loadingService.isLoading.set(false);

      this.notificationService.error('Please provide a valid email');
      return;
    }
    this.newsletterService.subscribeNewsletter(this.email).subscribe({
      next: (res) => {
        this.loadingService.isLoading.set(false);

        if (res.body?.message) {
          this.notificationService.success(res.body!.message);
        }
      },
      error: (err: HttpErrorResponse) => {
        this.loadingService.isLoading.set(false);

        if (err.error.error) {
          this.notificationService.error(err.error.error);
          return;
        }

        this.notificationService.error(`Could not complete request: Code ${err.status}`);
      },
    });
  }
}
