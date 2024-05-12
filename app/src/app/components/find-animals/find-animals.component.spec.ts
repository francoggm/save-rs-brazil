import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FindAnimalsComponent } from './find-animals.component';

describe('FindAnimalsComponent', () => {
  let component: FindAnimalsComponent;
  let fixture: ComponentFixture<FindAnimalsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FindAnimalsComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(FindAnimalsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
