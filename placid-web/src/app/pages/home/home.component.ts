import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import SoundScape from '@/shared/models/soundscape';
import fetchSoundscapes from '@/shared/services/sounds';

@Component({
  standalone: true,
  imports: [CommonModule],
  selector: 'app-home-page',
  templateUrl: 'home.html',
})
export class HomePageComponent {}
