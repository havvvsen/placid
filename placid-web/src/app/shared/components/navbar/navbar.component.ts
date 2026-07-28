import AuthService from '@/services/authservice';
import { Component, inject } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-navbar-component',
  templateUrl: 'navbar.html',
  styleUrls: ["navbar.css"],
  imports: [RouterLink],
})
export class NavbarComponent {
  authService = inject(AuthService)
  isLoggedIn: boolean = false

  constructor() {
    let token = localStorage.getItem("token")
    console.log(token)

    if (token == "" || token == null || token == undefined) {
      return
    }

    this.isLoggedIn = true
  }
}
