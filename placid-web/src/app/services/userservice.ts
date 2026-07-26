import { Environment } from '@/shared/constants/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject, Service } from '@angular/core';

@Service({
  autoProvided: true,
})
class UserService {
  private http = inject(HttpClient);
  private token

  constructor() {
    let token = localStorage.getItem("token")
    this.token = token
  }

  public getUser(email: string) {
    const headers = new HttpHeaders().set("Authorization", `Bearer ${this.token?.toString()}`)

    let body = {
      email: email
    }

    return this.http.post(Environment.endpoints.user, body, { headers: headers, observe: 'response' });
  }

  public deleteAccount(email: string, password: string) {
    const headers = new HttpHeaders().set("Authorization", `Bearer ${this.token?.toString()}`)

    let body = {
      email: email,
      password: password
    }

    return this.http.post(Environment.endpoints.deleteAccount, body, { headers: headers, observe: 'response' });
  }
}

export default UserService;
