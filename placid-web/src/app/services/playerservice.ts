import { inject, Injectable, signal } from '@angular/core';
import { Track } from '@/shared/models/track';
import { Environment } from '@/shared/constants/environment';
import { NotificationService } from './notificationservice';

@Injectable({
  providedIn: 'root',
})
export class PlayerService {
  notificationService = inject(NotificationService);
  public currentTrack = signal<Track | null>(null);
  public isPlaying = signal<boolean>(false);
  public trackList = signal<Track[]>([]);

  private audio: HTMLAudioElement | null = null;

  public setTrackList(tracks: Track[]) {
    this.trackList.set(tracks);
    if (!this.currentTrack() && tracks.length > 0) {
      this.currentTrack.set(tracks[0]);
    }
  }

  public playTrack(track: Track) {
    if (this.currentTrack()?.id === track.id && this.audio) {
      this.togglePlayStatus();
      return;
    }

    if (this.audio) {
      this.audio.pause();
      this.audio.onplay = null;
      this.audio.onpause = null;
    }

    this.currentTrack.set(track);
    this.audio = new Audio(`${Environment.trackServerBaseUrl}/${track?.audioUrl}`);
    this.audio.loop = true;

    this.audio.onplay = () => this.isPlaying.set(true);
    this.audio.onpause = () => this.isPlaying.set(false);

    this.audio
      .play()
      .then(() => {
        this.isPlaying.set(true);
      })
      .catch((err) => {
        if (err instanceof DOMException && err.name === 'AbortError') {
          return;
        }

        this.notificationService.error(err);
        this.isPlaying.set(false);
      });
  }

  public togglePlayStatus() {
    const current = this.currentTrack();
    if (!current) {
      const tracks = this.trackList();
      if (tracks.length > 0) {
        this.playTrack(tracks[0]);
      }
      return;
    }

    if (!this.audio) {
      this.audio = new Audio(current.audioUrl);
      this.audio.onplay = () => this.isPlaying.set(true);
      this.audio.onpause = () => this.isPlaying.set(false);
    }

    if (this.isPlaying()) {
      this.audio.pause();
    } else {
      this.audio
        .play()
        .then(() => {
          this.isPlaying.set(true);
        })
        .catch((err) => {
          if (err instanceof DOMException && err.name === 'AbortError') {
            return;
          }
          this.notificationService.error(err);
          this.isPlaying.set(false);
        });
    }
  }

  public pause() {
    if (this.audio && this.isPlaying()) {
      this.audio.pause();
    }
  }
}

export default PlayerService;
