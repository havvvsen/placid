import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterLink } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import AuthService from '@/services/authservice';


@Component({
  standalone: true,
  imports: [CommonModule, RouterLink, FormsModule],
  selector: 'app-register-page',
  templateUrl: 'register.html',
})
export class RegisterPageComponent {
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
      alert(e);
      return;
    }


    this.authService.registerUser(this.email, this.password).subscribe({
      next: (res) => {
        this.isLoading = false;
        this.router.navigateByUrl('/login');
        alert(res.body?.message);
      },
      error: (err: HttpErrorResponse) => {
        this.isLoading = false;
        console.log(`Status: ${err.status} . Message: ${err.message} . Url: ${err.url}`)
        alert(err.status);
      }
    })
  }
}
