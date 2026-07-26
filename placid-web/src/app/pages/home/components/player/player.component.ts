import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { PlayerService } from '@/services/playerservice';
import { Environment } from '@/shared/constants/environment';

@Component({
  selector: 'home-player-component',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: 'player.html',
  styleUrls: ['player.css'],
})
export class PlayerComponent {
  public playerService = inject(PlayerService);
  env = Environment

  formatTime(seconds: number): string {
    if (isNaN(seconds) || seconds < 0) return '0:00';
    const mins = Math.floor(seconds / 60);
    const secs = Math.floor(seconds % 60);
    return `${mins}:${secs < 10 ? '0' : ''}${secs}`;
  }

  togglePlayStatus() {
    this.playerService.togglePlayStatus();
  }

  playNext() {
    this.playerService.playNext();
  }

  playPrevious() {
    this.playerService.playPrevious();
  }

  onSeek(event: Event) {
    const input = event.target as HTMLInputElement;
    const value = parseFloat(input.value);
    this.playerService.seek(value);
  }

  onVolumeChange(event: Event) {
    const input = event.target as HTMLInputElement;
    const value = parseFloat(input.value);
    this.playerService.setVolume(value);
  }
}

