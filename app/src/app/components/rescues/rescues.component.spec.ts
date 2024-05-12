import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RescuesComponent } from './rescues.component';

describe('RescuesComponent', () => {
  let component: RescuesComponent;
  let fixture: ComponentFixture<RescuesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RescuesComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RescuesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
