import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '@/shared/components/navbar/navbar.component';
import { FooterComponent } from '@/shared/components/footer/footer.component';
import { Router, RouterLink } from '@angular/router';

@Component({
  standalone: true,
  imports: [CommonModule, NavbarComponent, FooterComponent, RouterLink],
  selector: 'app-landing-page',
  templateUrl: 'landing.html',
})
export class LandingPageComponent implements OnInit {
  private router = inject(Router);

  ngOnInit() {
    let token = localStorage.getItem('token');

    if (token) {
      this.router.navigateByUrl('/home');
    }
  }
}
