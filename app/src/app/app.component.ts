import { Component } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { RouterOutlet } from '@angular/router';
import { NavbarComponent } from './components/navbar/navbar.component';
import { Path } from './interfaces/path';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, NavbarComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'Save RS';

  paths: Path[] = [
    {
      path: "/posts",
      placeholder: "Todos",
      icon: "apps"
    },
    {
      path: "/resgates",
      placeholder: "Resgates",
      icon: "houseboat"
    },
    {
      path: "/animais-perdidos",
      placeholder: "Animais",
      icon: "pets"
    }
  ]

  public constructor(private titleService : Title){
    this.titleService.setTitle(this.title);
  }
}
