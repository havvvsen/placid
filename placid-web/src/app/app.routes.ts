import { Routes } from '@angular/router';
import { HomePageComponent } from './pages/home/home.component';
import { LoginPageComponent } from './pages/login/login.component';
import { RegisterPageComponent } from './pages/register/register.component';
import { LandingPageComponent } from './pages/landing/landing.component';
import { PlayerComponent } from './pages/player/player.component';

export const routes: Routes = [
  {
    path: '',
    component: LandingPageComponent,
  },
  {
    path: 'login',
    component: LoginPageComponent,
  },
  {
    path: 'register',
    component: RegisterPageComponent,
  },
  {
    path: 'home',
    component: HomePageComponent,
  },
  {
    path: 'player',
    component: PlayerComponent,
  },
];
