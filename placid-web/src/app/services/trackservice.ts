import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Environment } from '../shared/constants/environment';
import { Track } from '@/shared/models/track';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class TrackService {
  private http = inject(HttpClient);

  public getTrackList(): Observable<Track[]> {
    const token = localStorage.getItem('token');
    let headers = new HttpHeaders();
    if (token) {
      headers = headers.set('Authorization', `Bearer ${token.trim()}`);
    }

    return this.http.get<Track[]>(`${Environment.apiBaseUrl}/${Environment.endpoints.tracks}`, {
      headers,
    });
  }
}
