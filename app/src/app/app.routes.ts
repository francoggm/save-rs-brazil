import { Routes } from '@angular/router';
import { AllPostsComponent } from './components/all-posts/all-posts.component';
import { RescuesComponent } from './components/rescues/rescues.component';
import { FindAnimalsComponent } from './components/find-animals/find-animals.component';

export const routes: Routes = [
  {
    path: '',
    redirectTo: '/posts',
    pathMatch: 'full'
  },
  {
    path: 'posts',
    component: AllPostsComponent
  },
  {
    path: 'resgates',
    component: RescuesComponent
  },
  {
    path: 'animais-perdidos',
    component: FindAnimalsComponent
  }
];
