import Track from '@/shared/models/track';
import { inject, Service } from '@angular/core';
import { TrackList, TrackService } from './trackservice';

@Service({
  autoProvided: true,
})
class PlayerService {
  trackService = inject(TrackService);
  currentTrack: Track | undefined;
  trackList: TrackList | null;
  playing: boolean;

  constructor() {
    this.currentTrack = undefined;
    this.playing = false;

    this.trackList = this.trackService.getTrackList();
  }
}

export default PlayerService;
