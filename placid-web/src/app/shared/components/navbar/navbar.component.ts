import { Component } from '@angular/core';
import { ZardButtonComponent } from '@/shared/components/button';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-navbar-component',
  templateUrl: 'navbar.html',
  imports: [RouterLink],
})
export class NavbarComponent {}
