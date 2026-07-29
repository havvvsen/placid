import { TestBed } from '@angular/core/testing';
import { vi } from 'vitest';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import UserService from './userservice';
import { Environment } from '@/shared/constants/environment';

describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    vi.stubGlobal('localStorage', { getItem: vi.fn().mockReturnValue('mock-token') });
    
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [UserService]
    });

    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('getUser should send POST request with correct parameters and header', () => {
    service.getUser('test@example.com').subscribe();

    const req = httpMock.expectOne(Environment.endpoints.user);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ email: 'test@example.com' });
    expect(req.request.headers.get('Authorization')).toBe('Bearer mock-token');
    req.flush({});
  });

  it('deleteAccount should send POST request with correct parameters and header', () => {
    service.deleteAccount('test@example.com', 'password123').subscribe();

    const req = httpMock.expectOne(Environment.endpoints.deleteAccount);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ email: 'test@example.com', password: 'password123' });
    expect(req.request.headers.get('Authorization')).toBe('Bearer mock-token');
    req.flush({});
  });
});
