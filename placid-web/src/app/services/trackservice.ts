import { Mood } from '@/shared/models/mood';
import Track from '@/shared/models/track';
import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';

export class TrackList {
  focus: Track[] = [];
  relax: Track[] = [];
  sleep: Track[] = [];
}

@Service({
  autoProvided: true,
})
export class TrackService {
  private http = inject(HttpClient);
  private trackList: TrackList | null;

  constructor() {
    this.trackList = this.fetchTrackList();
  }

  public getTrackList(): TrackList | null {
    return this.trackList;
  }

  public fetchTrackList(): TrackList {
    // TODO: add fetch logic
    let tracks: Track[] = [
      {
        id: 1,
        name: 'Summer Tides',
        mood: 'focus',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 2,
        name: 'Weifer Heifer',
        mood: 'anxiety',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 3,
        name: 'Kneegor Meegro',
        mood: 'relax',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 4,
        name: 'CaseY Jones',
        mood: 'sleep',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 5,
        name: 'CaseY Jones',
        mood: 'focus',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 6,
        name: 'CaseY Jones',
        mood: 'focus',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 7,
        name: 'CaseY Jones',
        mood: 'focus',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
      {
        id: 8,
        name: 'CaseY Jones',
        mood: 'focus',
        audioUrl: 'http://localhost:8091/audio/stream',
        bgUrl: 'https://imgur/img.png',
      },
    ];

    let trackList: TrackList = new TrackList();

    tracks.map((track, _) => {
      switch (track.mood) {
        case 'relax':
          trackList.relax.push(track);
          break;
        case 'focus':
          trackList.focus.push(track);
          break;
        case 'sleep':
          trackList.sleep.push(track);
          break;
      }
    });

    console.log(trackList);
    return trackList;
  }
}
