
import { Environment } from '@/shared/constants/environment';
import { SubscribeNewsletterResponse } from '@/shared/models/response';
import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';

@Service({
  autoProvided: true,
})
class NewsletterService {
  private http = inject(HttpClient);

  public subscribeNewsletter(email: string) {

    let body = {
      email: email
    }

    return this.http.post<SubscribeNewsletterResponse>(`${Environment.apiBaseUrl}/${Environment.endpoints.subscribeNewsletter}`, body, { observe: 'response' });
  }

  public unsubscribeNewsletter(email: string) {
    let body = {
      email: email,
    }

    return this.http.delete(`${Environment.apiBaseUrl}/${Environment.endpoints.unSubscribeNewsletter}`, { body: body, observe: 'response' });
  }
}

export default NewsletterService;
