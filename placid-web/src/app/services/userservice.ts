import Endpoints from '@/shared/constants/endpoints';
import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';

@Service({
  autoProvided: true,
})
class UserService {
  private http = inject(HttpClient);
  private endPoints = new Endpoints();

  public registerUser(username: string, password: string) {
    let body = {
      username: username,
      password: password,
    };

    this.http.post(this.endPoints.register, body, {
      observe: 'response',
    });
  }

  public loginUser(username: string, password: string) {
    let body = {
      username: username,
      password: password,
    };

    this.http.post(this.endPoints.login, body, { observe: 'response' });
  }
}

export default UserService;
