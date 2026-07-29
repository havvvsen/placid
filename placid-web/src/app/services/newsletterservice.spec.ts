import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import NewsletterService from './newsletterservice';
import { Environment } from '@/shared/constants/environment';

describe('NewsletterService', () => {
  let service: NewsletterService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [NewsletterService]
    });
    service = TestBed.inject(NewsletterService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should subscribe to newsletter', () => {
    const email = 'test@example.com';
    service.subscribeNewsletter(email).subscribe(response => {
      expect(response.status).toBe(200);
    });

    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.subscribeNewsletter}`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ email });
    req.flush({}, { status: 200, statusText: 'OK' });
  });

  it('should unsubscribe from newsletter', () => {
    const email = 'test@example.com';
    service.unsubscribeNewsletter(email).subscribe(response => {
      expect(response.status).toBe(200);
    });

    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.unSubscribeNewsletter}`);
    expect(req.request.method).toBe('DELETE');
    expect(req.request.body).toEqual({ email });
    req.flush({}, { status: 200, statusText: 'OK' });
  });
});
