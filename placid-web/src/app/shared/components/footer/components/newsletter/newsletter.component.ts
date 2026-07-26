import NewsletterService from '@/services/newsletterservice';
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
  newsletterService = inject(NewsletterService)
  email = '';
  isLoading = false;

  onSubmit() {
    this.isLoading = true

    if (!this.email.includes('@') || !this.email.includes('.') || this.email.length < 4) {
      this.isLoading = false

      alert('Please provide a valid email');
      return;
    }
    this.newsletterService.subscribeNewsletter(this.email).subscribe({
      next: (res) => {
        this.isLoading = false;

        alert(res.body?.message);
      },
      error: (err: HttpErrorResponse) => {
        this.isLoading = false;
        alert(err.error.error);
      }
    })

  }

}
