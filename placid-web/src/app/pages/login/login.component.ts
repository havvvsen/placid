import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterLink } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import sanitizeCredentials from '@/shared/utils/sanitizer';

interface LoginResponse {
  message: string;
  token: string;
}

@Component({
  standalone: true,
  imports: [CommonModule, RouterLink, FormsModule],
  selector: 'app-login-page',
  templateUrl: 'login.html',
})
export class LoginPageComponent {
  private router = inject(Router);
  httpClient = inject(HttpClient);
  email = '';
  password = '';
  loading = false;

  onSubmit() {
    this.loading = true;

    try {
      sanitizeCredentials(this.email, this.password);
    } catch (e: any) {
      this.loading = false;
      alert(e);
      return;
    }

    const body = {
      email: this.email,
      password: this.password,
    };

    this.httpClient
      .post<LoginResponse>('http://localhost:3000/api/v1/auth/login', body, {
        observe: 'response',
      })
      .subscribe({
        next: (res) => {
          this.loading = false;

          if (res.status == 200) {
            if (res.body != null) {
              let token: string = res.body?.token;
              localStorage.setItem('token', token);
              this.router.navigateByUrl('/home');

              return;
            }
            alert('Invalid response');
          }
        },
        error: (err: HttpErrorResponse) => {
          this.loading = false;
          alert(err.error.error);
        },
      });
  }
}
