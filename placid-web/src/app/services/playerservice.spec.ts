import { TestBed } from '@angular/core/testing';
import { vi } from 'vitest';
import { PlayerService } from './playerservice';
import { NotificationService } from './notificationservice';
import { Track } from '@/shared/models/track';

describe('PlayerService', () => {
  let service: PlayerService;
  let mockNotificationService: any;
  let mockAudio: any;

  beforeEach(() => {
    mockNotificationService = {
      error: vi.fn()
    };

    mockAudio = {
      play: vi.fn().mockReturnValue(Promise.resolve()),
      pause: vi.fn(),
      loop: false
    };

    vi.spyOn(window, 'Audio').mockImplementation(function() { return mockAudio; } as any);

    TestBed.configureTestingModule({
      providers: [
        PlayerService,
        { provide: NotificationService, useValue: mockNotificationService }
      ]
    });

    service = TestBed.inject(PlayerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('setTrackList should update trackList and currentTrack if empty', () => {
    const tracks: Track[] = [{ id: 1, audioUrl: 'url1' } as Track, { id: 2, audioUrl: 'url2' } as Track];
    service.setTrackList(tracks);

    expect(service.trackList()).toEqual(tracks);
    expect(service.currentTrack()).toEqual(tracks[0]);
  });

  it('setTrackList should update trackList but not currentTrack if already set', () => {
    const tracks: Track[] = [{ id: 1, audioUrl: 'url1' } as Track];
    service.currentTrack.set({ id: 99, audioUrl: 'url99' } as Track);
    service.setTrackList(tracks);

    expect(service.trackList()).toEqual(tracks);
    expect(service.currentTrack()?.id).toBe(99);
  });

  it('playTrack should play a new track and set loop to true', async () => {
    const track: Track = { id: 1, audioUrl: 'url1' } as Track;
    service.playTrack(track);

    expect(service.currentTrack()).toEqual(track);
    expect(window.Audio).toHaveBeenCalled();
    expect(mockAudio.play).toHaveBeenCalled();
    expect(mockAudio.loop).toBe(true);
    
    // Simulate promise resolution
    await Promise.resolve();
    expect(service.isPlaying()).toBe(true);
  });

  it('pause should call audio pause and not change currentTrack', () => {
    const track: Track = { id: 1, audioUrl: 'url1' } as Track;
    service.playTrack(track);
    service.isPlaying.set(true);

    service.pause();
    expect(mockAudio.pause).toHaveBeenCalled();
    expect(service.currentTrack()).toEqual(track);
  });

  it('togglePlayStatus should play if track is paused', async () => {
    const track: Track = { id: 1, audioUrl: 'url1' } as Track;
    service.currentTrack.set(track);
    service.isPlaying.set(false);

    service.togglePlayStatus();
    expect(mockAudio.play).toHaveBeenCalled();
    
    await Promise.resolve();
    expect(service.isPlaying()).toBe(true);
  });

  it('togglePlayStatus should pause if track is playing', () => {
    const track: Track = { id: 1, audioUrl: 'url1' } as Track;
    service.playTrack(track); // Initializes this.audio
    service.isPlaying.set(true);

    service.togglePlayStatus();
    expect(mockAudio.pause).toHaveBeenCalled();
  });

  it('togglePlayStatus should play first track if no current track exists', () => {
    const tracks: Track[] = [{ id: 1, audioUrl: 'url1' } as Track];
    service.setTrackList(tracks);
    service.currentTrack.set(null); // Force null for this test case

    service.togglePlayStatus();
    expect(service.currentTrack()).toEqual(tracks[0]);
    expect(mockAudio.play).toHaveBeenCalled();
  });
});
