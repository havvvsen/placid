import { Environment } from '@/shared/constants/environment';
import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { RegisterResponse, LoginResponse } from '@/shared/models/response';
import InvalidInputError from '@/shared/exceptions/invalid_input';
import PlayerService from './playerservice';
import { Router } from '@angular/router';
import { NotificationService } from './notificationservice';

@Service({
  autoProvided: true,
})
export default class AuthService {
  private http = inject(HttpClient);
  private router = inject(Router);
  playerService = inject(PlayerService);
  notificationService = inject(NotificationService);

  registerUser(email: string, password: string) {
    let body = {
      email: email,
      password: password,
    };

    return this.http.post<RegisterResponse>(
      `${Environment.apiBaseUrl}/${Environment.endpoints.register}`,
      body,
      {
        observe: 'response',
      },
    );
  }

  loginUser(email: string, password: string) {
    let body = {
      email: email,
      password: password,
    };

    return this.http.post<LoginResponse>(
      `${Environment.apiBaseUrl}/${Environment.endpoints.login}`,
      body,
      { observe: 'response' },
    );
  }
  logoutUser() {
    if (this.playerService.isPlaying()) {
      this.playerService.togglePlayStatus();
    }

    localStorage.removeItem('token');
    this.notificationService.info('We look forward to seeing you again');

    this.router.navigateByUrl('/login', {
      replaceUrl: true,
    });
  }

  sanitizeCredentials(email: string, password: string) {
    if (!email.includes('@') || !email.includes('.') || email.length < 5) {
      throw new InvalidInputError('Please provide a valid email');
    }

    if (password.length < 6) {
      throw new InvalidInputError('Password cannot be less than 6 characters');
    }
  }
}
