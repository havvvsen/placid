import { Component, inject, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PlayerComponent } from './components/player/player.component';
import { TrackService } from '@/services/trackservice';
import { PlayerService } from '@/services/playerservice';
import { Track } from '@/shared/models/track';
import { TrackList } from '@/shared/models/tracklist';
import { HttpErrorResponse } from '@angular/common/http';
import { Environment } from '@/shared/constants/environment';
import { NavbarComponent } from '@/shared/components/navbar/navbar.component';
import TrackCardComponent from './components/trackcard/trackcard.component';
import { LoadingService } from '@/services/loadingservice';
import { NotificationService } from '@/services/notificationservice';

@Component({
  standalone: true,
  imports: [CommonModule, PlayerComponent, NavbarComponent, TrackCardComponent],
  selector: 'app-home-page',
  templateUrl: 'home.html',
})
export class HomePageComponent {
  private trackService = inject(TrackService);
  public playerService = inject(PlayerService);
  private cdr = inject(ChangeDetectorRef);
  loadingService = inject(LoadingService);
  notificationService = inject(NotificationService);
  trackList: TrackList | null = null;
  defaultTrack: Track | undefined = undefined;
  isLoading = false;
  env = Environment;

  constructor() {
    this.loadingService.isLoading.set(true);

    this.trackService.getTrackList().subscribe({
      next: (res: Track[]) => {
        this.loadingService.isLoading.set(false);

        const focusTracks = res.filter((t) => t.mood === 'focus');
        const relaxTracks = res.filter((t) => t.mood === 'relax');
        const sleepTracks = res.filter((t) => t.mood === 'sleep');

        this.trackList = {
          focus: focusTracks,
          relax: relaxTracks,
          sleep: sleepTracks,
        };

        this.defaultTrack = this.trackList.focus[0];
        this.playerService.setTrackList(res);
        this.cdr.detectChanges();
      },
      error: (err: HttpErrorResponse) => {
        this.loadingService.isLoading.set(false);

        if (err.error.error) {
          this.notificationService.error(err.error.error);

          return;
        }

        this.notificationService.error(`Could not complete request: Code ${err.status}`);
        this.cdr.detectChanges();
      },
    });
  }

  play(track: Track) {
    this.playerService.playTrack(track);
  }
}
