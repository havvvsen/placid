import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { httpResource } from '@angular/common/http';
import { PlayerComponent } from './components/player/player.component';

interface SoundScape {
  id: string;
  name: string;
  url: string;
}
enum Mood {
  focus = 'focus',
  relax = 'relax',
  sleep = 'sleep',
}
interface SoundScapesResponse {
  focus: SoundScape[];
  relax: SoundScape[];
  sleep: SoundScape[];
}

@Component({
  standalone: true,
  imports: [CommonModule, PlayerComponent],
  selector: 'app-home-page',
  templateUrl: 'home.html',
})
export class HomePageComponent {
  defaultSoundscape: SoundScape | null = null;

  moods = Object.values(Mood).map((value, index) => ({
    value,
    index,
  }));

  productsResource = httpResource<SoundScapesResponse>(() => ({
    url: 'http://localhost:3000/api/v1/soundscapes',
  }));
}
