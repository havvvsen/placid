import { HttpClient, HttpErrorResponse, HttpEvent, HttpResponse } from '@angular/common/http';
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
  email = '';
  loading = false;
  error = '';

  onSubmit() {
    const body = {
      email: this.email,
    };
    console.log(`Submitting ${this.email}`);

    if (!this.email.includes('@') || !this.email.includes('.') || this.email.length < 4) {
      alert('Please provide a valid email');
      return;
    }

    this.loading = true;
    this.error = '';

    this.http
      .post('http://localhost:3000/api/v1/join-newsletter', { body }, { observe: 'response' })
      .subscribe({
        next: (response) => {
          this.loading = false;
          this.email = '';

          if (response.status == 200) {
            alert('Thanks for subscribing!');
            return;
          }
        },

        error: (err: HttpErrorResponse) => {
          this.loading = false;
          if (err.status == 0) {
            alert('Server unavailable');
            return;
          }
          this.error = err.error.error;
          alert(this.error);
        },
      });
  }
}
