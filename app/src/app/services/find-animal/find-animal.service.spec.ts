import { TestBed } from '@angular/core/testing';

import { FindAnimalService } from './find-animal.service';

describe('FindAnimalService', () => {
  let service: FindAnimalService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FindAnimalService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
