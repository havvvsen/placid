import { Component } from '@angular/core';
import { FooterNewsLetterComponent } from './components/newsletter/newsletter.component';

@Component({
  selector: 'app-footer-component',
  standalone: true,
  imports: [FooterNewsLetterComponent],
  templateUrl: 'footer.html',
})
export class FooterComponent {

}
