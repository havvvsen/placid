import Track from '@/shared/models/track';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'home-track-card-component',
  templateUrl: 'trackcard.html',
})
export default class TrackCardComponent {
  @Input() track: Track;

  constructor(track: Track) {
    this.track = track;
  }
}
