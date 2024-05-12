import { Component, Input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIcon } from '@angular/material/icon';
import { MatToolbar, MatToolbarRow } from '@angular/material/toolbar';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field'
import { Router, RouterLink } from '@angular/router';
import { Path } from '../../interfaces/path';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [RouterLink, MatToolbar, MatToolbarRow, MatIcon, MatButtonModule, MatFormFieldModule, MatInputModule],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css'
})
export class NavbarComponent {
  @Input() title! : string;
  @Input() paths! : Path[];

  constructor(public router: Router){}
}
