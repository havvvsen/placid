import { TestBed } from '@angular/core/testing';
import { vi } from 'vitest';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { Router } from '@angular/router';
import AuthService from './authservice';
import PlayerService from './playerservice';
import { NotificationService } from './notificationservice';
import InvalidInputError from '@/shared/exceptions/invalid_input';
import { Environment } from '@/shared/constants/environment';

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;
  let mockRouter: any;
  let mockPlayerService: any;
  let mockNotificationService: any;

  beforeEach(() => {
    vi.stubGlobal('localStorage', { removeItem: vi.fn(), setItem: vi.fn(), getItem: vi.fn() });
    mockRouter = {
      navigateByUrl: vi.fn()
    };
    mockPlayerService = {
      isPlaying: vi.fn().mockReturnValue(true),
      togglePlayStatus: vi.fn()
    };
    mockNotificationService = {
      info: vi.fn()
    };

    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [
        AuthService,
        { provide: Router, useValue: mockRouter },
        { provide: PlayerService, useValue: mockPlayerService },
        { provide: NotificationService, useValue: mockNotificationService }
      ]
    });

    service = TestBed.inject(AuthService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('registerUser should send POST request', () => {
    service.registerUser('test@example.com', 'password123').subscribe();
    
    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.register}`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ email: 'test@example.com', password: 'password123' });
    req.flush({});
  });

  it('loginUser should send POST request', () => {
    service.loginUser('test@example.com', 'password123').subscribe();
    
    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.login}`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ email: 'test@example.com', password: 'password123' });
    req.flush({});
  });

  it('logoutUser should clear token, pause audio, notify, and redirect', () => {
    service.logoutUser();

    expect(mockPlayerService.isPlaying).toHaveBeenCalled();
    expect(mockPlayerService.togglePlayStatus).toHaveBeenCalled();
    expect(mockNotificationService.info).toHaveBeenCalledWith('We look forward to seeing you again');
    expect(mockRouter.navigateByUrl).toHaveBeenCalledWith('/login', { replaceUrl: true });
  });

  it('sanitizeCredentials should throw error on invalid email', () => {
    expect(() => service.sanitizeCredentials('invalid', 'password123')).toThrowError(InvalidInputError);
  });

  it('sanitizeCredentials should throw error on short password', () => {
    expect(() => service.sanitizeCredentials('test@example.com', '12345')).toThrowError(InvalidInputError);
  });

  it('sanitizeCredentials should not throw error on valid inputs', () => {
    expect(() => service.sanitizeCredentials('test@example.com', 'password123')).not.toThrow();
  });
});
