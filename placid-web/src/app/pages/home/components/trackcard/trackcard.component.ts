import PlayerService from '@/services/playerservice';
import { Track } from '@/shared/models/track';
import { Component, EventEmitter, inject, Input, Output } from '@angular/core';

@Component({
  selector: 'app-track-card-component',
  standalone: true,
  templateUrl: 'trackcard.html',
  host: {
    class: 'block w-full',
  },
})
export class TrackCardComponent {
  @Input({ required: true }) track!: Track;
  @Input({ required: true }) bannerServerBaseUrl!: string;
  @Output() play = new EventEmitter<Track>();

  playerService = inject(PlayerService);

  onPlay() {
    this.play.emit(this.track);
  }
}

export default TrackCardComponent;
