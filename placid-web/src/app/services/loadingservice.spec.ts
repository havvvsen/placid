import { TestBed } from '@angular/core/testing';
import { LoadingService } from './loadingservice';

describe('LoadingService', () => {
  let service: LoadingService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [LoadingService]
    });
    service = TestBed.inject(LoadingService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should initialize with isLoading as false', () => {
    expect(service.isLoading()).toBe(false);
  });

  it('show should set isLoading to true', () => {
    service.show();
    expect(service.isLoading()).toBe(true);
  });

  it('hide should set isLoading to false', () => {
    service.show(); // Set to true first
    service.hide();
    expect(service.isLoading()).toBe(false);
  });
});
