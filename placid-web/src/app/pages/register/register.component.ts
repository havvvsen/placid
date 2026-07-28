import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import AuthService from '@/services/authservice';
import { LoadingService } from '@/services/loadingservice';
import { NotificationService } from '@/services/notificationservice';
import InvalidInputError from '@/shared/exceptions/invalid_input';

@Component({
  standalone: true,
  imports: [CommonModule, FormsModule],
  selector: 'app-register-page',
  templateUrl: 'register.html',
})
export class RegisterPageComponent {
  private router = inject(Router);
  private authService = inject(AuthService);
  loadingService = inject(LoadingService);
  notificationService = inject(NotificationService);
  email = '';
  password = '';
  hidePassword = true;

  onSubmit() {
    this.loadingService.isLoading.set(true);

    try {
      this.authService.sanitizeCredentials(this.email, this.password);
    } catch (e: any) {
      this.loadingService.isLoading.set(false);

      if (e instanceof InvalidInputError) {
        this.notificationService.warning(e.message);

        return;
      }

      this.notificationService.error(e);
      return;
    }

    this.authService.registerUser(this.email, this.password).subscribe({
      next: (res) => {
        this.loadingService.isLoading.set(false);
        this.router.navigateByUrl('/login');
        alert(res.body?.message);
      },
      error: (err: HttpErrorResponse) => {
        this.loadingService.isLoading.set(false);
        this.notificationService.error(err.message);
      },
    });
  }
}
