import { TestBed } from '@angular/core/testing';
import { vi } from 'vitest';
import { NotificationService } from './notificationservice';
import { MatSnackBar } from '@angular/material/snack-bar';

describe('NotificationService', () => {
  let service: NotificationService;
  let snackBarSpy: any;

  beforeEach(() => {
    const spy = { open: vi.fn() };

    TestBed.configureTestingModule({
      providers: [
        NotificationService,
        { provide: MatSnackBar, useValue: spy }
      ]
    });
    service = TestBed.inject(NotificationService);
    snackBarSpy = TestBed.inject(MatSnackBar) as any;
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('success should open snackbar with correct parameters', () => {
    service.success('Success message');
    expect(snackBarSpy.open).toHaveBeenCalledWith('Success message', 'Close', {
      duration: 3000,
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      panelClass: ['success-snackbar'],
    });
  });

  it('error should open snackbar with correct parameters', () => {
    service.error('Error message');
    expect(snackBarSpy.open).toHaveBeenCalledWith('Error message', 'Close', {
      duration: 3000,
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      panelClass: ['error-snackbar'],
    });
  });

  it('warning should open snackbar with correct parameters', () => {
    service.warning('Warning message');
    expect(snackBarSpy.open).toHaveBeenCalledWith('Warning message', 'Close', {
      duration: 4000,
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      panelClass: ['warning-snackbar'],
    });
  });

  it('info should open snackbar with correct parameters', () => {
    service.info('Info message');
    expect(snackBarSpy.open).toHaveBeenCalledWith('Info message', 'Close', {
      duration: 3000,
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      panelClass: ['info-snackbar'],
    });
  });
});
