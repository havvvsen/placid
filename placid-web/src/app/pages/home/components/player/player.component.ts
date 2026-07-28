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
  env = Environment;

  togglePlayStatus() {
    this.playerService.togglePlayStatus();
  }
}
