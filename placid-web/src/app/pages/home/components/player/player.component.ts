import Track from '@/shared/models/track';
import { Component } from '@angular/core';

@Component({
  selector: 'home-player-component',
  templateUrl: 'player.html',
  styleUrls: ['player.css'],
})
export class PlayerComponent {
  isPlaying = false;
  currentTrack: Track | null = null;
  audio: HTMLAudioElement | null = null;

  constructor() {}

  playTrack() {
    try {
      this.audio = new Audio(this.currentTrack?.audioUrl);
      this.audio.play();
    } catch {
      alert('Failed to play soundscape');
    }
  }

  togglePlayStatus() {
    if (!this.currentTrack) {
      alert('Please select a track first');
      return;
    }

    if (this.isPlaying) {
      this.audio?.pause();
    } else {
      this.audio?.play();
    }

    this.isPlaying = !this.isPlaying;
  }
}
