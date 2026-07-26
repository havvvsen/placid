import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterLink } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import AuthService from '@/services/authservice';


@Component({
  standalone: true,
  imports: [CommonModule, RouterLink, FormsModule],
  selector: 'app-login-page',
  templateUrl: 'login.html',
})
export class LoginPageComponent {
  private router = inject(Router);
  private authService = inject(AuthService)
  email = '';
  password = '';
  isLoading = false;

  onSubmit() {
    this.isLoading = true;

    try {
      this.authService.sanitizeCredentials(this.email, this.password);
    } catch (e: any) {
      this.isLoading = false;
      alert(e);
      return;
    }

    this.authService.loginUser(this.email, this.password).subscribe({
      next: (res) => {
        this.isLoading = false;

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
        this.isLoading = false;
        alert(err.error.error);
      }
    })

  }
}
