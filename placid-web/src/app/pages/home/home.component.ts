import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { httpResource } from '@angular/common/http';
import { PlayerComponent } from './components/player/player.component';
import { TrackList, TrackService } from '@/services/trackservice';
import Track from '@/shared/models/track';

// interface TrackResponse {
//   focus: Track[];
//   relax: Track[];
//   sleep: Track[];
// }

@Component({
  standalone: true,
  imports: [CommonModule, PlayerComponent],
  selector: 'app-home-page',
  templateUrl: 'home.html',
})
export class HomePageComponent {
  soundscapeService = inject(TrackService);
  trackList: TrackList | null;
  defaultTrack: Track | undefined = undefined;

  // productsResource = httpResource<SoundScapesResponse>(() => ({
  //   url: 'http://localhost:3000/api/v1/soundscapes',
  // }));

  constructor() {
    this.trackList = this.soundscapeService.getTrackList();
    this.defaultTrack = this.trackList?.focus?.[0];
  }

  printTrackList() {
    console.log(this.trackList);
  }
}
