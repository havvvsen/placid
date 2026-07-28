import { Injectable, signal } from '@angular/core';
import { Track } from '@/shared/models/track';
import { Environment } from '@/shared/constants/environment';

@Injectable({
  providedIn: 'root',
})
export class PlayerService {
  public currentTrack = signal<Track | null>(null);
  public isPlaying = signal<boolean>(false);
  public currentTime = signal<number>(0);
  public duration = signal<number>(0);
  public volume = signal<number>(1);
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
      this.audio.ontimeupdate = null;
      this.audio.onloadedmetadata = null;
      this.audio.onended = null;
    }

    this.currentTrack.set(track);
    this.audio = new Audio(`${Environment.trackServerBaseUrl}/${track?.audioUrl}`);
    this.audio.loop = true;
    this.audio.volume = this.volume();

    this.audio.onplay = () => this.isPlaying.set(true);
    this.audio.onpause = () => this.isPlaying.set(false);
    this.audio.ontimeupdate = () => this.currentTime.set(this.audio?.currentTime || 0);
    this.audio.onloadedmetadata = () => this.duration.set(this.audio?.duration || 0);
    this.audio.onended = () => {
      this.isPlaying.set(false);
      this.playNext();
    };

    this.audio
      .play()
      .then(() => {
        this.isPlaying.set(true);
      })
      .catch((err) => {
        console.log(`${Environment.trackServerBaseUrl}/${track?.audioUrl}`);
        console.error('Playback error:', err);
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
      this.audio.volume = this.volume();
      this.audio.onplay = () => this.isPlaying.set(true);
      this.audio.onpause = () => this.isPlaying.set(false);
      this.audio.ontimeupdate = () => this.currentTime.set(this.audio?.currentTime || 0);
      this.audio.onloadedmetadata = () => this.duration.set(this.audio?.duration || 0);
      this.audio.onended = () => {
        this.isPlaying.set(false);
        this.playNext();
      };
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
          console.error('Playback error:', err);
          this.isPlaying.set(false);
        });
    }
  }

  public pause() {
    if (this.audio && this.isPlaying()) {
      this.audio.pause();
    }
  }

  public seek(seconds: number) {
    if (this.audio) {
      this.audio.currentTime = seconds;
      this.currentTime.set(seconds);
    }
  }

  public setVolume(volume: number) {
    const clamped = Math.max(0, Math.min(1, volume));
    this.volume.set(clamped);
    if (this.audio) {
      this.audio.volume = clamped;
    }
  }

  public playNext() {
    const tracks = this.trackList();
    if (tracks.length === 0) return;
    const current = this.currentTrack();
    if (!current) {
      this.playTrack(tracks[0]);
      return;
    }
    const index = tracks.findIndex((t) => t.id === current.id);
    const nextIndex = (index + 1) % tracks.length;
    this.playTrack(tracks[nextIndex]);
  }

  public playPrevious() {
    const tracks = this.trackList();
    if (tracks.length === 0) return;
    const current = this.currentTrack();
    if (!current) {
      this.playTrack(tracks[0]);
      return;
    }
    const index = tracks.findIndex((t) => t.id === current.id);
    const prevIndex = (index - 1 + tracks.length) % tracks.length;
    this.playTrack(tracks[prevIndex]);
  }
}

export default PlayerService;
