import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterLink } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import sanitizeCredentials from '@/shared/utils/sanitizer';

interface RegisterResponse {
  message: string;
}

@Component({
  standalone: true,
  imports: [CommonModule, RouterLink, FormsModule],
  selector: 'app-register-page',
  templateUrl: 'register.html',
})
export class RegisterPageComponent {
  private router = inject(Router);
  private httpClient = inject(HttpClient);
  email = '';
  password = '';
  loading = false;

  onSubmit() {
    this.loading = true;

    try {
      sanitizeCredentials(this.email, this.password);
    } catch (e: any) {
      alert(e);
      return;
    }

    const body = {
      email: this.email,
      password: this.password,
    };

    console.log(body);

    this.httpClient
      .post<RegisterResponse>('http://localhost:3000/api/v1/auth/register', body, {
        observe: 'response',
      })
      .subscribe({
        next: (res) => {
          this.loading = false;
          this.router.navigateByUrl('/login');
          alert(res.body?.message);
        },
        error: (err: HttpErrorResponse) => {
          this.loading = false;
          alert(err.error.error);
        },
      });
  }
}
