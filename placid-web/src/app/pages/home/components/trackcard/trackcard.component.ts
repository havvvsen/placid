import { Track } from '@/shared/models/track';
import { Component, Input, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PlayerService } from '@/services/playerservice';

@Component({
  selector: 'home-track-card-component',
  standalone: true,
  imports: [CommonModule],
  templateUrl: 'trackcard.html',
})
export class TrackCardComponent {
  @Input() track!: Track;
  public playerService = inject(PlayerService);

  get isPlaying(): boolean {
    return this.playerService.currentTrack()?.id === this.track?.id && this.playerService.isPlaying();
  }

  playTrack() {
    if (this.track) {
      this.playerService.playTrack(this.track);
    }
  }
}

export default TrackCardComponent;

