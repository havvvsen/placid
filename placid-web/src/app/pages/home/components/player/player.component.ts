import SoundScape from '@/shared/models/soundscape';
import { Component } from '@angular/core';

@Component({
  selector: 'home-player-component',
  templateUrl: 'player.html',
  styleUrls: ['player.css'],
})
export class PlayerComponent {
  isPlaying = false;
  currentSoundscape: SoundScape | null = null;
  audio: HTMLAudioElement | null = null;

  constructor() {}

  playSoundscape() {
    try {
      this.audio = new Audio(this.currentSoundscape?.audioUrl);
      this.audio.play();
    } catch {
      alert('Failed to play soundscape');
    }
  }

  togglePlayStatus() {
    if (!this.currentSoundscape) {
      alert('Please select a soundscape first');
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
