import { TestBed } from '@angular/core/testing';
import { vi } from 'vitest';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TrackService } from './trackservice';
import { Environment } from '../shared/constants/environment';
import { Track } from '@/shared/models/track';

describe('TrackService', () => {
  let service: TrackService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [TrackService]
    });

    service = TestBed.inject(TrackService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('getTrackList should include Authorization header if token exists', () => {
    vi.stubGlobal('localStorage', { getItem: vi.fn().mockReturnValue('mock-token') });

    service.getTrackList().subscribe();

    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.tracks}`);
    expect(req.request.method).toBe('GET');
    expect(req.request.headers.get('Authorization')).toBe('Bearer mock-token');
    req.flush([]);
  });

  it('getTrackList should not include Authorization header if token is absent', () => {
    vi.stubGlobal('localStorage', { getItem: vi.fn().mockReturnValue(null) });

    service.getTrackList().subscribe();

    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.tracks}`);
    expect(req.request.method).toBe('GET');
    expect(req.request.headers.has('Authorization')).toBe(false);
    req.flush([]);
  });

  it('getTrackList should return list of tracks', () => {
    vi.stubGlobal('localStorage', { getItem: vi.fn().mockReturnValue(null) });
    const mockTracks: Track[] = [{ id: 1, title: 'Track 1', audioUrl: 'url1' } as any];

    service.getTrackList().subscribe(tracks => {
      expect(tracks).toEqual(mockTracks);
    });

    const req = httpMock.expectOne(`${Environment.apiBaseUrl}/${Environment.endpoints.tracks}`);
    req.flush(mockTracks);
  });
});
