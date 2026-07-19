import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '@/shared/components/navbar/navbar.component';
import { FooterComponent } from '@/shared/components/footer/footer.component';

@Component({
  standalone: true,
  imports: [CommonModule, NavbarComponent, FooterComponent],
  selector: 'app-landing-page',
  templateUrl: 'landing.html',
})
export class LandingPageComponent {}
